package solution

import (
	"fmt"
	"strconv"
)

func Solution(N int) {
	for i := 1; i <= N; i++ {
		word := ""
		if i%2 == 0 {
			word += "Codility"
		}

		if i%3 == 0 {
			word += "Test"
		}

		if i%5 == 0 {
			word += "Coders"
		}

		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			word += strconv.Itoa(i)
		}

		fmt.Println(word)
	}
}
