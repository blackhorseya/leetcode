package main

func Solve(grid [][]byte) int {
	ret := 0

	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '1' {
				// start dfs
				sink(grid, i, j)
				ret++
			}
		}
	}

	return ret
}

func sink(grid [][]byte, i, j int) {
	grid[i][j] = '0'

	// up
	if i-1 >= 0 && grid[i-1][j] == '1' {
		sink(grid, i-1, j)
	}

	// down
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		sink(grid, i+1, j)
	}

	// right
	if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
		sink(grid, i, j+1)
	}

	// left
	if j-1 >= 0 && grid[i][j-1] == '1' {
		sink(grid, i, j-1)
	}
}

func main() {

}
