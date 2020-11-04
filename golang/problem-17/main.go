package main

func Solve(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	return recursive("", digits)
}

var phone = map[byte][]byte{'1': nil,
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func recursive(prefix, digits string) (ret []string) {
	if digits == "" {
		return []string{prefix}
	}

	for _, c := range phone[digits[0]] {
		ret = append(ret, recursive(prefix + string(c), digits[1:])...)
	}

	return
}

func main() {

}
