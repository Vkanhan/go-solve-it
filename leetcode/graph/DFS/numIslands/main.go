package main 

import(
	"fmt"
)

func numIslands(grid [][]bool) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || visited[r][c] || !grid[r][c] {
			return 
		}

		visited[r][c] = true 

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] && !visited[r][c] {
				count++
				dfs(r, c)
			}
		}
	}
	return count 
}

func main() {
	grid := [][]bool{
		{true, true, false, false, false},
		{true, true, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true, true},
	}

	fmt.Println("Number of islands:", numIslands(grid))
}