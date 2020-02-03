package problem0207

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numCourses    int
	prerequisites [][]int
	ans           bool
}{

	{
		4,
		[][]int{
			[]int{1, 0},
			[]int{2, 1},
			[]int{2, 0},
			[]int{3, 0},
			[]int{3, 1},
			[]int{3, 2},
		},
		true,
	},

	{
		4,
		[][]int{
			[]int{1, 0},
			[]int{2, 1},
			[]int{2, 0},
			[]int{3, 0},
			[]int{3, 1},
			[]int{3, 2},
			[]int{2, 3},
		},
		false,
	},

	{
		2,
		[][]int{
			[]int{0, 1},
		},
		true,
	},

	{
		2,
		[][]int{
			[]int{1, 0},
		},
		true,
	},

	{
		2,
		[][]int{
			[]int{1, 0},
			[]int{0, 1},
		},
		false,
	},
	// 可以有多个 testcase
}

// TestCanFinish is a test
func TestCanFinish(t *testing.T) {

	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		// canFinish(tc.numCourses, tc.prerequisites)
		ast.Equal(tc.ans, canFinish(tc.numCourses, tc.prerequisites), "输入:%v", tc)
	}
}
