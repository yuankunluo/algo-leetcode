package problem0568

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCase = []struct {
	flights [][]int
	days    [][]int
	output  int
}{
	{
		[][]int{
			[]int{0, 1, 0},
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
		[][]int{
			[]int{0, 0, 7},
			[]int{2, 0, 0},
			[]int{7, 7, 7},
		},
		7,
	},
	{
		[][]int{
			[]int{0, 1, 0},
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
		[][]int{
			[]int{0, 0, 7},
			[]int{2, 7, 7},
			[]int{7, 7, 7},
		},
		16,
	},
	{
		[][]int{
			[]int{0, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 0},
		},
		[][]int{
			[]int{1, 3, 1},
			[]int{6, 0, 3},
			[]int{3, 3, 3},
		},
		12,
	},
	{
		[][]int{
			[]int{0, 0, 0},
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
		[][]int{
			[]int{1, 1, 1},
			[]int{7, 7, 7},
			[]int{7, 7, 7},
		},
		3,
	},
}

func TestMaxVacationDays(t *testing.T) {
	ast := assert.New(t)
	for i, tc := range testCase {
		fmt.Println("case", i)
		ast.Equal(tc.output, maxVacationDays(tc.flights, tc.days))
	}
}
