package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCommandLineArgs(t *testing.T) {
	argument := "1 2 3 4 5 6 7 8 9 command"
	exp, command := ParseCommandLineArgs(argument)
	assert.Equal(t, "1 2 3 4 5 6 7 8 9", exp)
	assert.Equal(t, "command", command)

	argument = ""
	exp, command = ParseCommandLineArgs(argument)
	assert.Equal(t, "", exp)
	assert.Equal(t, "", command)

	argument = "command"
	exp, command = ParseCommandLineArgs(argument)
	assert.Equal(t, "", exp)
	assert.Equal(t, "command", command)
}
