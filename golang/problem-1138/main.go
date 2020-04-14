package main

// Define a coordinate structure
type Point struct {
	X int32
	Y int32
}

// Get pointer of Point by letter
func NewPointByLetter(letter int32) *Point {
	pos := letter - 'a'
	return &Point{
		X: pos % 5,
		Y: pos / 5,
	}
}

// Get the path of "from" to "to" in the fourth quadrant
func (from *Point) WalkTo(to *Point) string {
	var result = ""

	for from.Y > to.Y {
		result += "U"
		from.Y--
	}
	
	for from.X > to.X {
		result += "L"
		from.X--
	}
	
	for from.Y < to.Y {
		result += "D"
		from.Y++
	}
	
	for from.X < to.X {
		result += "R"
		from.X++
	}

	return result + "!"
}

func Solve(target string) string {
	var result = ""

	var current = &Point{}
	for _, letter := range target {
		goal := NewPointByLetter(letter)
		result += current.WalkTo(goal)
	}

	return result
}

func main() {

}
