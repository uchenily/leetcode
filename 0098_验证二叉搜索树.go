package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 思路,中序遍历, 将结果保存在一个数组(切片)中, 如果数组是递增的, 则是二叉搜索树.
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	treeArray := midTraverse(root)
	for i := 0; i < len(treeArray)-1; i++ {
		if treeArray[i] >= treeArray[i+1] {
			return false
		}
	}
	return true
}

func midTraverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	left := midTraverse(root.Left)
	result = append(left, root.Val)
	right := midTraverse(root.Right)
	result = append(result, right...)
	return result
}
