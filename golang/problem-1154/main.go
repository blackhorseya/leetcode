package main

import (
	"time"
)

func Solve(input string) int {
	layout := "2006-01-02"
	if date, err := time.Parse(layout, input); err != nil {
		return 0
	} else {
		return date.YearDay()
	}
}

func main() {
	Solve("")
}
