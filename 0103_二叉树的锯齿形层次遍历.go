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
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	depth := getDepth(root)
	for i := 0; i < depth; i++ {
		levelN := getLevelN(root, i)
		if i%2 == 0 {
			for j := 0; j < len(levelN)/2; j++ {
				levelN[j], levelN[len(levelN)-1-j] = levelN[len(levelN)-1-j], levelN[j]
			}
		}
		result = append(result, levelN)
		fmt.Println(getLevelN(root, i))
	}
	return result
}

// 获取第N层的数据集合
func getLevelN(root *TreeNode, n int) []int {
	if root == nil {
		return []int{}
	}
	if n == 0 {
		return []int{root.Val}
	}
	result := make([]int, 0)
	left := getLevelN(root.Left, n-1)
	right := getLevelN(root.Right, n-1)
	result = append(result, right...)
	result = append(result, left...)
	return result
}

// 获取二叉树深度
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	result := left + 1
	if right > left {
		result = right + 1
	}
	return result
}
