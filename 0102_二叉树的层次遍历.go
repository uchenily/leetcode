package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	depth := treeDepth(root)
	result := make([][]int, depth)
	for i := 0; i < depth; i++ {
		result[i] = getLevelN(root, i)
	}
	return result
}

// 返回二叉树的深度
func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeDepth(root.Left)
	right := treeDepth(root.Right)
	result := left + 1
	if left < right {
		result = right + 1
	}
	return result
}

// 获取第n层的数据(0<=n<depth), 返回一个切片
// 利用递归, getLevelN(root, n) = getLevelN(root.Left, n-1) + getLevelN(root.Right, n-1)
func getLevelN(root *TreeNode, n int) []int {
	if root == nil {
		return []int{}
	}
	if n == 0 {
		return []int{root.Val}
	}
	result := make([]int, 0)
	result = append(result, getLevelN(root.Left, n-1)...)
	result = append(result, getLevelN(root.Right, n-1)...)
	return result
}
