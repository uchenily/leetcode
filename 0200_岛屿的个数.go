package leetcode

func numIslands(grid [][]byte) int {
	row := len(grid)
	if len(grid) == 0 {
		return 0
	}
	col := len(grid[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	result := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				result++
				dfs(grid, visited, i, j)
			}
		}
	}
	// dfs(grid, visited, 0, 0)
	// fmt.Println(visited)
	return result
}

func dfs(grid [][]byte, visited [][]bool, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if !visited[i][j] && grid[i][j] == '1' {
		// fmt.Println(i+1, j+1) // 搜索路径
		visited[i][j] = true
		dfs(grid, visited, i, j-1)
		dfs(grid, visited, i, j+1)
		dfs(grid, visited, i-1, j)
		dfs(grid, visited, i+1, j)
	}
}
