package problem0404

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// getHeight 返回树的高度。
//
// 递归计算树的高度是典型的递归运算。
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

// max 返回最大值。
//
// 一个辅助函数。
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printTree(root *TreeNode) [][]string {
	// 计算输出数组的规模
	h := getHeight(root)
	// 宽度（最底层的宽度）， 2^n - 1
	w := (1 << h) - 1
	// 构造答案二维数组（slice）
	ans := make([][]string, h)
	for i := range ans {
		ans[i] = make([]string, w)
	}
	fill(root, &ans, 0, 0, w-1)
	return ans
}

func fill(root *TreeNode, ans *[][]string, h, l, r int) {
	if root == nil {
		return
	}
	mid := (l + r) / 2
	*ans[h][mid] = strconv.Itoa(root.Val)
	fill(root.Left, ans, h+1, l, mid-1)
	fill(root.Right, ans, h+1, mid+1, r)
}
