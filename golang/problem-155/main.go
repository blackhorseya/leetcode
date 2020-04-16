package main

import "math"

type Node struct {
	Value int
	Min   int
}

type MinStack struct {
	Nodes []*Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := this.GetMin()
	if x < min {
		min = x
	}

	this.Nodes = append(this.Nodes, &Node{
		Value: x,
		Min:   min,
	})
}

func (this *MinStack) Pop() {
	if len(this.Nodes) == 0 {
		return
	}

	this.Nodes = this.Nodes[:len(this.Nodes)-1]
}

func (this *MinStack) Top() int {
	if len(this.Nodes) == 0 {
		return 0
	}

	return this.Nodes[len(this.Nodes)-1].Value
}

func (this *MinStack) GetMin() int {
	if len(this.Nodes) == 0 {
		return math.MaxInt32
	}

	return this.Nodes[len(this.Nodes)-1].Min
}

func main() {

}
