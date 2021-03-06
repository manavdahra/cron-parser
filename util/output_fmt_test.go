package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOutputFmt(t *testing.T) {
	expected := `minute        %s
hour          %s
day of month  %s
month         %s
day of week   %s
command       %s`
	assert.Equal(t, expected, GenerateOutputFmt())
}
