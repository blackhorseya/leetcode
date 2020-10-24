package main

func Solve(rowIndex int) []int {
	ret := make([]int, rowIndex+1)

	for index := 0; index < len(ret); index++ {
		if index == 0 {
			ret[index] = 1
		} else if index == 1 {
			ret[index] = rowIndex
		} else {
			ret[index] = (ret[index-1] * (rowIndex - index + 1)) / index
		}
	}

	return ret
}

func main() {

}
