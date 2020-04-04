package main

import (
	"strconv"
)

func Solve(a string, b string) string {
	var alen = len(a)
	var blen = len(b)

	var str = ""
	var posa = alen - 1
	var posb = blen - 1
	var carry uint8
	for (posa >= 0 && posb >= 0) || carry >= 1 {
		var aval, bval int
		if posa >= 0 {
			aval, _ = strconv.Atoi(string(a[posa]))
		}

		if posb >= 0 {
			bval, _ = strconv.Atoi(string(b[posb]))
		}

		var val = aval + bval + int(carry)
		str = strconv.Itoa(val % 2) + str

		if val >= 2 {
			carry = 1
		} else {
			carry = 0
		}

		posa--
		posb--
	}

	if posa >= 0 {
		str = string(a[0:posa + 1]) + str
	}

	if posb >= 0 {
		str = string(b[0:posb + 1]) + str
	}

	return str
}

func main() {
	
}
