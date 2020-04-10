package main

import "fmt"

type MinStack struct {
	Stack []int
	Min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	if len(this.Min) == 0 || x <= this.GetMin() {
		this.Min = append(this.Min, x)
	}
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.Min = this.Min[:len(this.Min)-1]
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

func main() {
	minStack := &MinStack{[]int{}, []int{}}
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin()
	minStack.Pop()
	minStack.Top()
	minStack.GetMin()
	fmt.Println(minStack)
}
