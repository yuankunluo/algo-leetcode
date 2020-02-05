package problem0671

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	return dfs(root, root.Val)
}

// dfs is to find the second smallest element in this tree.
// s is the current smallest we have.
func dfs(root *TreeNode, s int) int {

	// If root is nil, that means we can't find anything,
	// we return -1
	if root == nil {
		return -1
	}

	// If the current tree value is bigger than the smallest
	// element we hat, that means we have found the second smallest
	// one in the tree, since all subtrees is greater or equal to this
	// value.
	if root.Val > s {
		return root.Val
	}

	sLeft := dfs(root.Left, s)
	sRight := dfs(root.Right, s)

	if sLeft == -1 {
		return sRight
	}

	if sRight == -1 {
		return sLeft
	}

	return min(sLeft, sRight)

}

// min a helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
