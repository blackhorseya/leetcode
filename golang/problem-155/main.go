package main

import "math"

type MinStack struct {
	Values   []int
	Minimums []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Values = append(this.Values, x)

	min := this.GetMin()
	if x < min {
		min = x
	}

	this.Minimums = append(this.Minimums, min)
}

func (this *MinStack) Pop() {
	length := len(this.Values)
	if length == 0 {
		return
	}

	this.Values = this.Values[:length-1]
	this.Minimums = this.Minimums[:length-1]
}

func (this *MinStack) Top() int {
	if len(this.Values) == 0 {
		return 0
	}

	return this.Values[len(this.Values)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Minimums) == 0 {
		return math.MaxInt32
	}

	return this.Minimums[len(this.Minimums)-1]
}

func main() {

}
