package problem0108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	root := &TreeNode{
		Val:  nums[mid],
		Left: sortedArrayToBST(nums[:mid]),
	}
	if mid+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}
