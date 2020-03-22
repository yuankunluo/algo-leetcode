package problem0063

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  [][]int
	output int
}{
	// {
	// 	input: [][]int{
	// 		[]int{
	// 			1,
	// 		},
	// 	},
	// 	output: 0,
	// },
	{
		input: [][]int{
			[]int{
				0, 0,
			},
		},
		output: 0,
	},
}

func TestUniquePathsWithObstacles(t *testing.T) {

	ast := assert.New(t)

	for _, v := range testCases {
		ast.Equal(v.output, uniquePathsWithObstacles(v.input))
	}
}
