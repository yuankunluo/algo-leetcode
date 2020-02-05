package problem1239

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input  []string
	output int
}{
	{
		[]string{"un", "iq", "ue"},
		4,
	},
	{
		[]string{"cha", "r", "act", "ers"},
		6,
	},
	{
		[]string{"abcdefghijklmnopqrstuvwxyz"},
		26,
	},
}

func TestMaxLength(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range testCases {
		ast.Equal(tc.output, maxLength(tc.input))
	}
}
