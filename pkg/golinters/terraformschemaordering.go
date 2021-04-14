package golinters
//
//import (
//	"github.com/mlin-aviatrix/Terraform-Schema-Ordering/pkg/analyzer"
//	"golang.org/x/tools/go/analysis"
//)
//
//func NewTerraformSchemaOrdering() *goanalysis.Linter {
//	return goanalysis.NewLinter(
//		"terraformschemaordering",
//		"Checks resource schema field order for terraform",
//		[]*analysis.Analyzer{analyzer.Analyzer},
//		nil,
//		).WithLoadMode(goanalysis.LoadModeSyntax)
//}