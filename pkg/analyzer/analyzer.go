package analyzer

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"reflect"
	"sort"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "terraformschemachecker",
	Doc:      "Checks resource schema field order for terraform",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var orderMap = make(map[string]int)

func run(pass *analysis.Pass) (interface{}, error) {
	// Setup reverse mapping from field name to index in array
	for i, elem := range ordering {
		orderMap[elem] = i
	}

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		// Suppress any errors
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()

		// Start with a function declaration
		funcDecl := node.(*ast.FuncDecl)
		funcName := fmt.Sprintf("%s", funcDecl.Name)
		// Only consider functions with "resource" or "dataSource" prefix
		if !strings.HasPrefix(funcName, "resource") && !strings.HasPrefix(funcName, "dataSource") {
			return
		}

		// Function should only have 1 return value of type *schema.Resource
		if len(funcDecl.Type.Results.List) != 1 {
			return
		}
		field := funcDecl.Type.Results.List[0]
		if reflect.TypeOf(field.Type) != reflect.TypeOf((*ast.StarExpr)(nil)) {
			return
		}
		starExpr := field.Type.(*ast.StarExpr)
		if reflect.TypeOf(starExpr.X) != reflect.TypeOf((*ast.SelectorExpr)(nil)) {
			return
		}
		selectorExpr := starExpr.X.(*ast.SelectorExpr)
		if fmt.Sprintf("%s", selectorExpr.X) != "schema" || fmt.Sprintf("%s", selectorExpr.Sel) != "Resource" {
			return
		}

		// Look for return statement
		var retStmt *ast.ReturnStmt
		for _, expr := range funcDecl.Body.List {
			if reflect.TypeOf(expr) == reflect.TypeOf((*ast.ReturnStmt)(nil)) {
				retStmt = expr.(*ast.ReturnStmt)
			}
		}
		// If no return statement found, return
		if retStmt == nil {
			return
		}

		// Function should only return 1 thing
		if len(retStmt.Results) != 1 {
			return
		}
		result := retStmt.Results[0]

		if reflect.TypeOf(result) != reflect.TypeOf((*ast.UnaryExpr)(nil)) {
			return
		}

		unary := result.(*ast.UnaryExpr)
		compositeLit := unary.X.(*ast.CompositeLit)

		parseResource(compositeLit, pass)
	})

	return nil, nil
}

func parseResource(compositeLit *ast.CompositeLit, pass *analysis.Pass) {
	selExpr := compositeLit.Type.(*ast.SelectorExpr)
	// Only consider schema.Resource
	if fmt.Sprintf("%s", selExpr.X) != "schema" || fmt.Sprintf("%s", selExpr.Sel) != "Resource" {
		return
	}
	var schemaKeyValueExpr *ast.KeyValueExpr
	for _, expr := range compositeLit.Elts {
		keyValueExpr := expr.(*ast.KeyValueExpr)
		// Only consider schema.Schema definition
		if fmt.Sprintf("%s", keyValueExpr.Key) == "Schema" {
			schemaKeyValueExpr = keyValueExpr
			break
		}
	}
	// Return if schema.Schema is not found
	if schemaKeyValueExpr == nil {
		return
	}

	// For each attribute, check the order of the fields
	attributes := schemaKeyValueExpr.Value.(*ast.CompositeLit)
	for _, expr := range attributes.Elts {
		field := expr.(*ast.KeyValueExpr)
		checkFieldOrder(field, pass)
	}
	return
}

func checkFieldOrder(field *ast.KeyValueExpr, pass *analysis.Pass) {
	fieldDef := field.Value.(*ast.CompositeLit)
	sorted := make([]int, len(fieldDef.Elts))

	for i, expr := range fieldDef.Elts {
		keyValueExpr := expr.(*ast.KeyValueExpr)
		key := fmt.Sprintf("%s", keyValueExpr.Key)
		v, ok := orderMap[key]
		if ok {
			sorted[i] = v
		} else {
			pass.Reportf(keyValueExpr.Pos(), "found new field name: %s", key)
			// Prevent other fields from conflicting with this field
			sorted[i] = len(ordering) + 1
		}
	}
	sort.Ints(sorted)

	indexes := make([]int, len(fieldDef.Elts))

	for i, expr := range fieldDef.Elts {
		keyValueExpr := expr.(*ast.KeyValueExpr)
		key := fmt.Sprintf("%s", keyValueExpr.Key)
		v, ok := orderMap[key]
		if ok {
			indexes[i] = v
		} else {
			pass.Reportf(keyValueExpr.Pos(), "found new field name: %s", key)
			// Prevent other fields from conflicting with this field
			indexes[i] = len(ordering) + 1
		}

		index := getWrongIndex(indexes, v, i)
		if index == - 1 {
			continue
		} else {
			name2 := ordering[indexes[index]]
			pass.Reportf(keyValueExpr.Pos(), "%s should come before %s", key, name2)
		}

		// Handle nested resource definitions
		if key == "Elem" {
			unary := keyValueExpr.Value.(*ast.UnaryExpr)
			compositeLit := unary.X.(*ast.CompositeLit)
			parseResource(compositeLit, pass)
		}
	}
}

func getWrongIndex(slice []int, val int, end int) int {
	smallestValue := 0
	smallestIndex := -1
	for i := 0; i < end; i++ {
		v := slice[i]
		if v < val {
			continue
		}
		if smallestIndex == -1 || v < smallestValue {
			smallestValue = v
			smallestIndex = i
		}
	}
	return smallestIndex
}

func sliceIndex(slice []int, val int) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}