package genutil

import (
	"fmt"
	"go/ast"
	"go/token"
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
	assert.Len(nodes, 4)
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

	falseFilter := func(node ast.Node) bool {
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
		structDecl, ok := genDecl.Specs[0].(*ast.ValueSpec)
		if !ok {
			return ok
		}
		fmt.Printf("Struct: %v\n", structDecl)
		// for _, spec := range genDecl.Specs {
		// 	fmt.Printf("     GenDecl spec: %v\n", spec)
		// }
		return ok
	}

	nodes, err := FilterAstNodesFromPatterns(falseFilter, "../_example")
	if err != nil {
		log.Error(err)
		assert.Error(err)
	}
	assert.Len(nodes, 1)
}
