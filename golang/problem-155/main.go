package main

import "math"

type MinStack struct {
	Values []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Values = append(this.Values, x)
}

func (this *MinStack) Pop() {
	length := len(this.Values)
	this.Values = this.Values[:length-1]
}

func (this *MinStack) Top() int {
	if len(this.Values) == 0 {
		return 0
	}

	return this.Values[len(this.Values)-1]
}

func (this *MinStack) GetMin() int {
	result := math.MaxInt64
	for _, value := range this.Values {
		result = int(math.Min(float64(result), float64(value)))
	}

	return result
}

func main() {

}
