func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				searchIslands(grid, &visited, i, j)
				res++
			}
		}
	}
	return res
}

func searchIslands(grid [][]byte, visited *[][]bool, i, j int) {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[i])-1 || (*visited)[i][j] {
		return
	}

	(*visited)[i][j] = true
	if grid[i][j] == '1' {
		searchIslands(grid, visited, i-1, j)
		searchIslands(grid, visited, i, j-1)
		searchIslands(grid, visited, i+1, j)
		searchIslands(grid, visited, i, j+1)
	}
}
