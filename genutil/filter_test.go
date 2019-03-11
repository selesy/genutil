package genutil

import (
	"go/ast"
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
		return true //TODO - filter some of the incoming nodes
	}

	nodes, err := FilterAstNodesFromPatterns(structArrayFilter, "../_example")
	if err != nil {
		log.Error(err)
		assert.Error(err)
	}
	assert.Len(nodes, 8) //TODO - see above
}
