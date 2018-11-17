package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	result := inorder(root)
	return result[k-1]
}

// 中序遍历
func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	result = append(result, inorder(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorder(root.Right)...)
	return result
}
