package problem0606

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	s := strconv.Itoa(t.Val)
	l := tree2str(t.Left)
	r := tree2str(t.Right)

	// case: leaf
	if t.Left == nil && t.Right == nil {
		return s
	}

	// case: right node is nil
	if t.Right == nil {
		return s + "(" + l + ")"
	}

	// general case : s(l)(r)
	return s + "(" + l + ")" + "(" + r + ")"

}
