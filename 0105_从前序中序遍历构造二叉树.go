package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// []preorder [0][..left..][..right..]
// []inorder [..left..][index][..right..]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		fmt.Println("Val:", preorder[0])
		return &TreeNode{Val: preorder[0]}
	}
	index := 0
	// 找到中序遍历中根结点位置
	for i := 0; i < len(preorder); i++ {
		if preorder[0] == inorder[i] {
			index = i
			break
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}
