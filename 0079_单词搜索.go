package leetcode

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	if row == 0 || col == 0 {
		return false
	}
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, visited, []byte(word), i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited [][]bool, word []byte, i, j int) bool {
	var result bool
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	if len(word) == 1 && board[i][j] == word[0] && !visited[i][j] { // 需要是未访问过的结点
		return true
	}
	if visited[i][j] {
		return false
	} else if board[i][j] == word[0] { // 尚未访问
		visited[i][j] = true // 标记
		result = dfs(board, visited, word[1:], i, j-1) ||
			dfs(board, visited, word[1:], i, j+1) ||
			dfs(board, visited, word[1:], i-1, j) ||
			dfs(board, visited, word[1:], i+1, j)
		visited[i][j] = false // 复原
	}
	return result
}
