package problem0110

// TreeNode is node
//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced is to check if a tree is balanced
//
func isBalanced(root *TreeNode) bool {
	// nil tree is always balanced
	if root == nil {
		return true
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	return (abs(leftHeight-rightHeight) <= 1) && isBalanced(root.Left) && isBalanced(root.Right)
}

// height return a tree's height
//
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

// abs computes the absolute value of a number.
//
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return 0 - a
}

// max a helper function to get max from 2 ints
//
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
