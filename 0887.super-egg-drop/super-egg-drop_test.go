package problem0887

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	K      int
	N      int
	output int
}{
	{
		K:      1,
		N:      2,
		output: 2,
	},
}

func TestSuperEggDrop(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range testCases {
		ast.Equal(tc.output, superEggDrop(tc.K, tc.N))
	}
}
