package problem0102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	// Create a queue.
	queue := make([]*TreeNode, 0)
	// Enqueue the root note
	queue = append(queue, root)

	for len(queue) > 0 {
		// Get the nodeCount in the queue,
		// which is also the nodeCount in current level
		nodeCount := len(queue)
		level := make([]int, 0)
		// Dequeue all nodes from the current level
		// Enqueue all nodes of the next level
		for nodeCount > 0 {
			node := queue[0]
			level = append(level, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			nodeCount--
		}
		result = append(result, level)
	}
	return result
}
