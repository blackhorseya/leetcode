package main

import (
	"fmt"
)

func Solve(target string) string {
	var result = ""

	var flagRow, flagCol int
	for _, letter := range target {
		pos := int(letter) - int('a')
		row := pos / 5
		col := pos % 5

		for flagRow > row {
			result += "U"
			flagRow--
		}

		for flagCol > col {
			result += "L"
			flagCol--
		}

		for flagRow < row {
			result += "D"
			flagRow++
		}

		for flagCol < col {
			result += "R"
			flagCol++
		}

		result += "!"
		fmt.Println(pos, row, col, result)
	}

	return result
}

func main() {

}
