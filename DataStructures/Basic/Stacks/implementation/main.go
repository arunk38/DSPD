package main

/*
 * Basic implementaion of stack using golang with standard stack functionalities
 */

import "fmt"

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Returns true if stack is empty, else false.
func (this *Stack) isEmpty() bool {
	if this.length == 0 {
		return true
	}
	return false
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
	fmt.Print("Pushing ", value, ", Stack ->")
	this.Print()
}

// Print the contents of current stack
func (this *Stack) Print() {
	fmt.Print("[")
	n := this.top
	for n != nil {
		fmt.Print(" ", n.value)
		n = n.prev
	}
	fmt.Print(" ]\n")
}

func main() {
	s := New()

	if s.isEmpty() {
		fmt.Println("Stack is empty currently")
	}

	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(9)
	s.Print()
	fmt.Println("Pop the first element: ", s.Pop())
	fmt.Println("Top element in stack currently: ", s.Peek())
	fmt.Println("Is stack empty: ", s.isEmpty())
	s.Print()
}
