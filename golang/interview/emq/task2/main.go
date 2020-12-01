package main

func Solution(S string) string {
	charset := []rune(S)
	flag := 0

	for flag < len(charset)-1 {
		if charset[flag] == 67 && charset[flag+1] == 68 {
			charset = append(charset[:flag], charset[flag+2:]...)
			flag = 0
			continue
		}

		if charset[flag] == 68 && charset[flag+1] == 67 {
			charset = append(charset[:flag], charset[flag+2:]...)
			flag = 0
			continue
		}

		if charset[flag] == 66 && charset[flag+1] == 65 {
			charset = append(charset[:flag], charset[flag+2:]...)
			flag = 0
			continue
		}

		if charset[flag] == 65 && charset[flag+1] == 66 {
			charset = append(charset[:flag], charset[flag+2:]...)
			flag = 0
			continue
		}

		flag++
	}

	return string(charset)
}

func main() {

}
