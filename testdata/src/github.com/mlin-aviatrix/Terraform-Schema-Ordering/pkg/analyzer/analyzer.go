package analyzer

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"reflect"

	//"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "terraformschemachecker",
	Doc: "Checks resource schema field order",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var ordering = []string{"Type", "Elem", "Required", "Optional", "Computed", "Default", "ForceNew", "Sensitive", "ValidateFunc", "Description", "ConflictsWith", "RequiredWith", "Deprecated"}
var orderMap = make(map[string]int, len(ordering))

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ReturnStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		retStmt := node.(*ast.ReturnStmt)
		if len(retStmt.Results) != 1 {
			return
		}
		result := retStmt.Results[0]

		if reflect.TypeOf(result) != reflect.TypeOf((*ast.UnaryExpr)(nil)) {
			return
		}

		unary := result.(*ast.UnaryExpr)
		compositeLit := unary.X.(*ast.CompositeLit)

		selExpr := compositeLit.Type.(*ast.SelectorExpr)
		// Only consider schema.Resource
		if fmt.Sprintf("%s", selExpr.X) != "schema" || fmt.Sprintf("%s", selExpr.Sel) != "Resource"  {
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

		// Setup reverse mapping from field name to index in array
		for i, elem := range ordering {
			orderMap[elem] = i
		}

		// For each attribute, check the order of the fields
		attributes := schemaKeyValueExpr.Value.(*ast.CompositeLit)
		for _, expr := range attributes.Elts {
			field := expr.(*ast.KeyValueExpr)
			checkFieldOrder(field, pass)
		}
		return
	})

	return nil, nil
}

func checkFieldOrder(field *ast.KeyValueExpr, pass *analysis.Pass) {
	fieldDef := field.Value.(*ast.CompositeLit)
	indexes := make([]int, len(fieldDef.Elts))

	for i, expr := range fieldDef.Elts {
		keyValueExpr := expr.(*ast.KeyValueExpr)
		key := fmt.Sprintf("%s", keyValueExpr.Key)
		var ok bool
		indexes[i], ok = orderMap[key]
		if !ok {
			pass.Reportf(keyValueExpr.Pos(), "found new field name: %s", key)
			// Prevent other fields from conflicting with this field
			indexes[i] = len(ordering) + 1
		}
		if i == 0 {
			continue
		}
		if indexes[i] < indexes[i - 1] {
			name1 := ordering[indexes[i]]
			name2 := ordering[indexes[i-1]]

			pass.Reportf(keyValueExpr.Pos(), "%s should come before %s in resource schema definition", name1, name2)
		}
	}
}