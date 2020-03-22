package problem0692

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  []string
	k      int
	output []string
}{
	{
		input: []string{"i",
			"love",
			"leetcode",
			"i",
			"love",
			"coding"},
		k: 2,
		output: []string{
			"i", "love",
		},
	},
	{
		input: []string{
			"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is",
		},
		k: 4,
		output: []string{
			"the", "is", "sunny", "day",
		},
	},
}

func TestTopKFreq(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range testCases {
		ast.Equal(tc.output, topKFrequent(tc.input, tc.k))
	}
}
