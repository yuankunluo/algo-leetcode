/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Leetcode: Medium
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	
}

// 
func inorderTraverse(root *TreeNode, ans []int) {
	if root == nil {
		return
	}
	inorderTraverse(root.Left, ans)
	ans = append(ans, root.Val)
	inorderTraverse(root.Left, ans)
}


func main(){
	
}