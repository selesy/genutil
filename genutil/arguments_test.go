package genutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyArgs(t *testing.T) {
	var argsIn []string
	assert.Empty(t, nonFlagArgs(argsIn))
}

func TestOnlyFlagArgs(t *testing.T) {
	argsIn := []string{"flag0", "flag1", "flag2"}
	assert.Len(t, nonFlagArgs(argsIn), 3)
}

func TestFlagsExcluded(t *testing.T) {
	argsIn := []string{"arg0", "-flag0", "arg1", "-flag1", "arg2"}
	assert.Len(t, nonFlagArgs(argsIn), 3)
}
