package main

import "fmt"

type MyStack struct {
	Stack []rune
}

func (this *MyStack) Push(value rune) {
	this.Stack = append(this.Stack, value)
}

func (this *MyStack) Pop() {
	if !this.IsEmpty() {
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
}

func (this *MyStack) IsEmpty() bool {
	return len(this.Stack) == 0
}

func output(s string) string {
	stack := &MyStack{[]rune{}}

	for _, v := range s {
		if v == '#' {
			stack.Pop()
		} else {
			stack.Push(v)
		}
	}
	return string(stack.Stack)
}

func backspaceCompare(S string, T string) bool {
	return output(S) == output(T)
}

func main() {
	S := "ab#c"
	T := "ab#c"
	fmt.Println(backspaceCompare(S, T))
}
