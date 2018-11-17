package leetcode

func rotate(matrix [][]int) {
	length := len(matrix)
	// 先沿主对角线旋转
	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 再按照中轴(Y)旋转
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-1-j] = matrix[i][length-1-j], matrix[i][j]
		}
	}
}
