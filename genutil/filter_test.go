package genutil

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNoneFiltered(t *testing.T) {
	assert := assert.New(t)

	trueFilter := func(node ast.Node) bool {
		return true
	}

	nodes, err := FilterAstNodesFromPatterns(trueFilter, "../_example")
	if err != nil {
		log.Error(err)
		assert.Error(err)
	}
	assert.Len(nodes, 8)
}

func TestAllFiltered(t *testing.T) {
	assert := assert.New(t)

	falseFilter := func(node ast.Node) bool {
		return false
	}

	nodes, err := FilterAstNodesFromPatterns(falseFilter, "../_example")
	if err != nil {
		log.Error(err)
		assert.Error(err)
	}
	assert.Len(nodes, 0)
}

func TestStructFiltered(t *testing.T) {
	assert := assert.New(t)

	structArrayFilter := func(node ast.Node) bool {
		fmt.Println("-----")

		//_, ok := node.(*ast.ValueSpec)
		genDecl, ok := node.(*ast.GenDecl)
		if !ok {
			return ok
		}
		if genDecl.Tok != token.VAR {
			return false
		}

		fmt.Printf("     GenDecl: %v\n", genDecl)
		fmt.Printf("     GenDecl CG: %v\n", genDecl.Doc)

		if len(genDecl.Specs) != 1 {
			return false
		}
		valueSpec, ok := genDecl.Specs[0].(*ast.ValueSpec)
		if !ok {
			return ok
		}
		fmt.Printf("     Spec: %v\n", valueSpec)

		compLit, ok := valueSpec.Values[0].(*ast.CompositeLit)
		if !ok {
			return false
		}
		fmt.Println("     CompositeLit: ", compLit)
		fmt.Println("     CompositeLit type: ", compLit.Type)
		fmt.Println("     CompositeLit type: ", reflect.TypeOf(compLit.Type))
		arrayType, ok := compLit.Type.(*ast.ArrayType)
		if !ok {
			return false
		}
		fmt.Println("     ArrayType:", arrayType)
		fmt.Println("     ArrayType element:", arrayType.Elt)
		arrayStructType, ok := arrayType.Elt.(*ast.StructType)
		if !ok {
			return false
		}
		fmt.Println("     ", arrayStructType)
		for _, field := range arrayStructType.Fields.List {
			fmt.Println("     ArrayStruct field:", field)
		}

		for _, compLitElt := range compLit.Elts {
			fmt.Println("     CompositeLit element: ", compLitElt)
			fmt.Println("     CompositeLit element type: ", reflect.TypeOf(compLitElt))
			innerCompLit, ok := compLitElt.(*ast.CompositeLit)
			if !ok {
				return false
			}
			for _, innerElt := range innerCompLit.Elts {
				fmt.Println("     Inner element:", innerElt)
			}
		}

		return ok
	}

	nodes, err := FilterAstNodesFromPatterns(structArrayFilter, "../_example")
	if err != nil {
		log.Error(err)
		assert.Error(err)
	}
	assert.Len(nodes, 1)
}
