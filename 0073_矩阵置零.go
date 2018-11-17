package leetcode

func setZeroes(matrix [][]int) {
	// row col 分别存放存放是否删除该行该列的标志
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	// fmt.Println("row", row)
	// fmt.Println("col", col)
	for i := 0; i < len(matrix); i++ {
		if row[i] == true {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if col[j] == true {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
}
