package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"github.com/mlin-aviatrix/Terraform-Schema-Ordering/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}