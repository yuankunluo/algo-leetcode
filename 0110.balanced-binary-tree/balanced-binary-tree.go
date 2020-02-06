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
	if root == nil {
		return true
	}
	balanced := true
	height(root, &balanced)
	return balanced
}

// height return a tree's height
//
func height(root *TreeNode, balanced *bool) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left, balanced)
	rightHeight := height(root.Right, balanced)
	if abs(leftHeight-rightHeight) > 1 {
		*balanced = false
		return -1
	}
	return max(leftHeight, rightHeight) + 1
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
