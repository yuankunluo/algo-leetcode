package problem0113

// TreeNode represents a node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	// The answer to be returned.
	ans := make([][]int, 0)
	// The current answer for one path.
	// This slice store the temporary list of values
	// that sum up to be the sum we want.
	// It works like an accumulator.
	cur := make([]int, 0)
	findPathSum(root, sum, &cur, &ans)
	return ans
}

func findPathSum(root *TreeNode, sum int, cur *[]int, ans *[][]int) {
	// Empty Node do nothing.
	if root == nil {
		return
	}

	// Leave Node
	if root.Left == nil && root.Right == nil {
		// We find the answer
		if root.Val == sum {
			// First append the value into the current slice.
			*cur = append(*cur, root.Val)
			// Append the final accumulator's values into answer.
			// Watch out this time we need append a *Copy* of current slice
			copy := make([]int, 0)
			copy = append(copy, (*cur)...)
			*ans = append(*ans, copy)
			// Pop this value out
			*cur = (*cur)[:len(*cur)-1]
		}
		return
	}

	// Middle stage
	*cur = append(*cur, root.Val)
	newSum := sum - root.Val
	findPathSum(root.Left, newSum, cur, ans)
	findPathSum(root.Right, newSum, cur, ans)
	*cur = (*cur)[:len(*cur)-1]
}
