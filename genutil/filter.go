package genutil

import "go/ast"

type AstNodeFilter func(goFile ast.File) []ast.Node

func FilterAstNodesFromPkgs(filter AstNodeFilter, pkgs Pkgs) ([]ast.Node, error) {
	var nodes []ast.Node
	return nodes, nil
}

func FilterAstNodesFromPatterns(filter AstNodeFilter, patterns ...string) ([]ast.Node, error) {
	pkgs, err := PackagesFromPatterns(patterns...)
	if err != nil {
		return nil, err
	}

	return FilterAstNodesFromPkgs(filter, pkgs)
}

func FilterAstNodesFromArgs(filter AstNodeFilter) ([]ast.Node, error) {
	pkgs, err := PackagesFromArgs()
	if err != nil {
		return nil, err
	}

	return FilterAstNodesFromPkgs(filter, pkgs)
}
