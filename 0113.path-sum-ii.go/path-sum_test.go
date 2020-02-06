package problem0113

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTree = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	},
	Right: &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 13,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	},
}

var output = [][]int{
	[]int{5, 4, 11, 2},
	[]int{5, 8, 4, 5},
}

var sum = 22

func TestPathSum(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(output, pathSum(testTree, sum))
}
