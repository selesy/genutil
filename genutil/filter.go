package genutil

import (
	"go/ast"
	"go/parser"
)

type AstNodeFilter func(node ast.Node) bool

func FilterAstNodes(filter AstNodeFilter, unfiltered []ast.Node) (filtered []ast.Node, err error) {
	for _, node := range unfiltered {
		if filter(node) {
			filtered = append(filtered, node)
		}
	}
	return
}
func FilterAstNodesFromArgs(filter AstNodeFilter) ([]ast.Node, error) {
	pkgs, err := PackagesFromArgs()
	if err != nil {
		return nil, err
	}

	return FilterAstNodesFromPkgs(filter, pkgs)
}

func FilterAstNodesFromFile(filter AstNodeFilter, file *ast.File) ([]ast.Node, error) {
	var decls []ast.Node
	for _, decl := range file.Decls {
		decls = append(decls, decl)
	}
	return FilterAstNodes(filter, decls)
}

func FilterAstNodesFromPatterns(filter AstNodeFilter, patterns ...string) ([]ast.Node, error) {
	pkgs, err := PackagesFromPatterns(patterns...)
	if err != nil {
		return nil, err
	}

	return FilterAstNodesFromPkgs(filter, pkgs)
}

func FilterAstNodesFromPkgs(filter AstNodeFilter, pkgs Pkgs) (filtered []ast.Node, err error) {
	mode := parser.ParseComments
	for _, pkg := range pkgs.pkgs {
		for _, goFile := range pkg.GoFiles {
			file, err := parser.ParseFile(pkgs.FileSet(), goFile, nil, mode)
			if err != nil {
				return nil, err
			}
			nodes, err := FilterAstNodesFromFile(filter, file)
			if err != nil {
				return nil, err
			}
			filtered = append(filtered, nodes...)
		}
	}
	return
}
