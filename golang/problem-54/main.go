package main

func Solve(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	count := m * n
	row, col := 0, 0
	find := make([][]int, m)
	for index := range find {
		find[index] = make([]int, n)
	}

	var ret []int

	for count > 0 {
		ret = append(ret, matrix[row][col])
		find[row][col] = 1
		count--

		// right
		if col+1 < n && find[row][col+1] != 1 {
			// up
			if row-1 >= 0 && find[row-1][col] != 1 {
				row--
				continue
			}
			col++
			continue
		}

		// down
		if row+1 < m && find[row+1][col] != 1 {
			row++
			continue
		}

		// left
		if col-1 >= 0 && find[row][col-1] != 1 {
			col--
			continue
		}

		// up
		if row-1 >= 0 && find[row-1][col] != 1 {
			row--
			continue
		}
	}

	return ret
}

func main() {

}
