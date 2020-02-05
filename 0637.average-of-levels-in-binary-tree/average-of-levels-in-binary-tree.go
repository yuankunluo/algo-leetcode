package problem0637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	if root == nil {
		return result
	}

	// Create a queue
	queue := make([]*TreeNode, 0)
	// Enqueue the root node
	queue = append(queue, root)

	for len(queue) > 0 {
		nodeCount := len(queue)
		nodeNum := nodeCount
		level := make([]int, 0)
		for nodeCount > 0 {
			node := queue[0]
			level = append(level, node.Val)
			// Dequeue the visited node
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			nodeCount--
		}
		result = append(result, float64(sum(level))/float64(nodeNum))
	}
	return result
}

// sum Compute the sum of a slice
func sum(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}
