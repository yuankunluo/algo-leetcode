package problem200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  [][]byte
	output int
}{
	{
		input: [][]byte{
			[]byte{
				'1', '1', '1', '1', '0',
			},
			[]byte{
				'1', '1', '0', '1', '0',
			},
			[]byte{
				'1', '1', '0', '0', '0',
			},
			[]byte{
				'0', '0', '0', '0', '0',
			},
		},
		output: 1,
	},
}

func TestnumIsland(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range testCases {
		ast.Equal(numIslands(tc.input), tc.output)
	}
}
