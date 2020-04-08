package main

var (
	first  = [][]int{{1}}
	second = [][]int{{1}, {1, 1}}
)

func Solve(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	if numRows == 1 {
		return first
	}

	result := second
	for index := 2; index < numRows; index++ {
		last := result[len(result)-1]
		newRow := make([]int, len(last)+1)
		for j := 0; j < len(newRow); j++ {
			if j == 0 || j == len(newRow)-1 {
				newRow[j] = 1
				continue
			}
			newRow[j] = last[j] + last[j-1]
		}
		result = append(result, newRow)
	}

	return result
}

func main() {

}
