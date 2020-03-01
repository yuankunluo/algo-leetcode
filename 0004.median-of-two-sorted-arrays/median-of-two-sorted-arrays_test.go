package problem0004

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	n1     []int
	n2     []int
	median float64
}{
	{
		n1:     []int{1, 2},
		n2:     []int{3, 4},
		median: 2.5,
	},
}

func TestFindMeianSortedArrays(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range testCases {
		ans := findMedianSortedArrays(tc.n1, tc.n2)
		fmt.Println(ans)
		ast.Equal(tc.median, ans)
	}
}
