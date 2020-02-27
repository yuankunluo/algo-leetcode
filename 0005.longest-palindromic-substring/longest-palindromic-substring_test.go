package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	output string
}{
	{
		input:  "babad",
		output: "bab",
	},
}

func TestLongestPalindrom(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range testCases {
		ast.Equal(tc.output, longestPalindrome(tc.input))
	}
}
