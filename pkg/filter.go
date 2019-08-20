package genutil

import (
	"go/ast"
	"go/parser"
)

type AstNodeFilter func(node ast.Node) bool

type FilterMatch struct {
	File   *ast.File
	GoFile string
	Node   ast.Node
}

func FilterAstNodes(filter AstNodeFilter, unfiltered []ast.Node) (filtered []ast.Node) {
	for _, node := range unfiltered {
		if filter(node) {
			filtered = append(filtered, node)
		}
	}
	return
}
func FilterAstNodesFromArgs(filter AstNodeFilter) ([]FilterMatch, error) {
	pkgs, err := PackagesFromArgs()
	if err != nil {
		return nil, err
	}

	return FilterAstNodesFromPkgs(filter, pkgs)
}

func FilterAstNodesFromFile(filter AstNodeFilter, file *ast.File) (matches []FilterMatch) {
	var decls []ast.Node
	for _, decl := range file.Decls {
		decls = append(decls, decl)
	}
	for _, node := range FilterAstNodes(filter, decls) {
		matches = append(matches, FilterMatch{
			File: file,
			Node: node,
		})
	}
	return
}

func FilterAstNodesFromPatterns(filter AstNodeFilter, patterns ...string) ([]FilterMatch, error) {
	pkgs, err := PackagesFromPatterns(patterns...)
	if err != nil {
		return nil, err
	}
	return FilterAstNodesFromPkgs(filter, pkgs)
}

func FilterAstNodesFromPkgs(filter AstNodeFilter, pkgs Pkgs) (filtered []FilterMatch, err error) {
	mode := parser.ParseComments
	for _, pkg := range pkgs.pkgs {
		for _, goFile := range pkg.GoFiles {
			file, err := parser.ParseFile(pkgs.FileSet(), goFile, nil, mode)
			if err != nil {
				return nil, err
			}
			matches := FilterAstNodesFromFile(filter, file)
			for _, match := range matches {
				(&match).GoFile = goFile
			}
			filtered = append(filtered, matches...)
		}
	}
	return
}
