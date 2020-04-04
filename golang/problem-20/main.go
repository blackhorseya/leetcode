package main

func Solve(s string) bool {
	if s == "" {
		return true
	}

	var q []string
	for _, v := range s {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			q = append(q, string(v))
		}

		if len(q) == 0 {
			return false
		}

		last := q[len(q)-1]
		if string(v) == ")" || string(v) == "}" || string(v) == "]" {
			if last != "(" && string(v) == ")" {
				return false
			}
			if last != "{" && string(v) == "}" {
				return false
			}
			if last != "[" && string(v) == "]" {
				return false
			}
			q = q[:len(q)-1]
		}
	}

	if len(q) > 0 {
		return false
	}

	return true
}

func main() {
	
}
