package problem1319

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	n           int
	connections [][]int
	ans         int
}{
	{
		n:   4,
		ans: 1,
		connections: [][]int{
			[]int{0, 1},
			[]int{0, 2},
			[]int{1, 2},
		},
	},
}

func TestMakeConnected(t *testing.T) {

	ast := assert.New(t)

	for _, tc := range testCases {
		ast.Equal(tc.ans, makeConnected(tc.n, tc.connections))
	}
}
