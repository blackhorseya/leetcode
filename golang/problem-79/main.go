package main

func Solve(board [][]byte, word string) bool {
	pos := 0

	for i, row := range board {
		for j := range row {
			if board[i][j] == word[pos] {
				// start dfs
				if ret := dfs(board, i, j, pos, word); ret {
					return true
				}

			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j, pos int, word string) bool {
	// out of boundary
	if i < 0 || j < 0 {
		return false
	}

	// out of boundary
	if i >= len(board) || j >= len(board[0]) {
		return false
	}

	// already visited
	if board[i][j] == '*' {
		return false
	}

	// difference character
	if word[pos] != board[i][j] {
		return false
	}

	// the end position
	if pos == len(word)-1 {
		return true
	}

	pos++
	temp := board[i][j]
	board[i][j] = '*'

	ret := dfs(board, i-1, j, pos, word) || // up
		dfs(board, i, j+1, pos, word) || // right
		dfs(board, i+1, j, pos, word) || // down
		dfs(board, i, j-1, pos, word) // left

	board[i][j] = temp

	return ret
}

func main() {

}
