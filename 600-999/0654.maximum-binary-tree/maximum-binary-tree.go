package problem0654


// TreeNode is node
//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
		return nil
	}
	maxIndex := findMaxElementIndex(nums)
	root := &TreeNode{
		Val: nums[maxIndex],
	}
	root.Left = constructMaximumBinaryTree(nums[0:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

// findMaxElementIndex, find the maximum's index.
//
func findMaxElementIndex(nums []int) int {
	maxI := 0
	maxV := 0
	for i, v := range nums {
		if v > maxV {
			maxI = i
			maxV = v
		}
	}
	return maxI
}