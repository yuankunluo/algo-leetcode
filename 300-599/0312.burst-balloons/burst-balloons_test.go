package problem0312

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  []int
	output int
}{
	{
		input:  []int{3, 1, 5, 8},
		output: 167,
	},
}

func TestMaxCoins(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range testCases {
		ast.Equal(tc.output, maxCoins(tc.input))
	}
}
