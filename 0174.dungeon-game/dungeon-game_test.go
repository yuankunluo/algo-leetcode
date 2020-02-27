package problem0174

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input [][]int 
	output int
}{
	{
		input: [][]int{
			[]int{
				-2,
				-3,
				3,
			},
			{
				-5,
				-10,
				1,
			},
			{
				10,30,-5,
			},
		},
		output: 7,
	},
}


func TestCalculateMinimHP(t *testing.T){
	ast := assert.New(t)

	for _, tc := range testCases {
		ast.Equal(tc.output, calculateMinimumHP(tc.input))
	}
}