package problem0153

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCase = []struct {
	input []int
	ans   int
}{
	{
		input: []int{3, 1, 2},
		ans:   1,
	},
	{
		input: []int{4, 5, 6, 7, 0, 1, 2},
		ans:   0,
	},
}

func TestFindMim(t *testing.T) {

	ast := assert.New(t)

	for i, tc := range testCase {
		fmt.Printf("Test case %d with %v", i, tc)
		ast.Equal(tc.ans, findMin(tc.input))
	}
}
