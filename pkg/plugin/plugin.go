package main

import (
	"github.com/mlin-aviatrix/Terraform-Schema-Ordering/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}

var AnalyzerPlugin analyzerPlugin