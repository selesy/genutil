package test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	// "golang.org/x/tools/go/packages"
	genutil "github.com/selesy/go-genutil/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// None of these comments should be included in the ``Test`` variable
// below as there's a line break before the element's comment.
//directive: not-a-thing-1
//directive-1: not-a-thing-2

//directive: Thing-1
//directive-1: Thing-2
//directive-2: Cat-in-the-Hat
//directive-3: Billy
//directive-3: Sally
var Test string

func getGenDecls(decls []ast.Decl) []*ast.GenDecl {
	var genDecls []*ast.GenDecl
	for _, decl := range decls {
		if g, ok := decl.(*ast.GenDecl); ok {
			genDecls = append(genDecls, g)
		}
	}
	return genDecls
}

func getValueSpecs(specs []ast.Spec) []*ast.ValueSpec {
	var valueSpecs []*ast.ValueSpec
	for _, spec := range specs {
		if valueSpec, ok := spec.(*ast.ValueSpec); ok {
			valueSpecs = append(valueSpecs, valueSpec)
		}
	}
	return valueSpecs
}

func getValueSpecWithName(specs []ast.Spec, name string) (*ast.ValueSpec, bool) {
	for _, spec := range getValueSpecs(specs) {
		for _, ident := range spec.Names {
			if ident.Name == name {
				return spec, true
			}
		}
	}
	return nil, false
}

func getVarDecls(decls []ast.Decl) []*ast.GenDecl {
	var varDecls []*ast.GenDecl
	for _, decl := range getGenDecls(decls) {
		if decl.Tok == token.VAR {
			varDecls = append(varDecls, decl)
		}
	}
	return varDecls
}

func getVarDeclByName(decls []ast.Decl, name string) (*ast.GenDecl, bool) {
	for _, decl := range getVarDecls(decls) {
		if _, ok := getValueSpecWithName(decl.Specs, name); ok {
			return decl, true
		}
	}
	return nil, false
}

func TestCommentGroupDirectives(t *testing.T) {
	fset := token.NewFileSet()
	mode := parser.ParseComments
	file, err := parser.ParseFile(fset, "directive_test.go", nil, mode)
	require.NoError(t, err)

	decl, ok := getVarDeclByName(file.Decls, "Test")
	require.True(t, ok)
	directives, err := genutil.CommentGroupWithDefaultConfig(decl, "directive")
	assert.NoError(t, err)
	assert.Len(t, directives, 4)
}
