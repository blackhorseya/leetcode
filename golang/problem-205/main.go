package main

func Solve(s string, t string) bool {
	st := make(map[rune]rune)
	charset1 := []rune(s)
	charset2 := []rune(t)

	for index := 0; index < len(s); index++ {
		c1 := charset1[index]
		c2 := charset2[index]

		if value, ok := st[c1]; ok {
			if value != c2 {
				return false
			}
		} else {
			st[c1] = c2

			for k, v := range st {
				if v == c2 && k != c1 {
					return false
				}
			}
		}
	}

	return true
}

func main() {

}
