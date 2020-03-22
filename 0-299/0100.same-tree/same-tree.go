package problem0100

// TreeNode is node
//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both are nil
	if p == nil && q == nil {
		return true
	}

	// One is nil but the other is not
	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	// Recursion call on subtrees
	return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
