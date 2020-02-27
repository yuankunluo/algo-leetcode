package problem1278

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	s      string
	k      int
	output int
}{
	{
		s:      "abc",
		k:      2,
		output: 1,
	},
	{
		s:      "aabbc",
		k:      3,
		output: 0,
	},
}

func TestPalindromPartitions(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range testCases {
		ast.Equal(tc.output, palindromePartition(tc.s, tc.k))
	}
}
