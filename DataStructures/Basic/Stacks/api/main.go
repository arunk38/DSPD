package main

/*
 * Basic implementaion of stack using golang with standard stack functionalities
 */

import (
	"fmt"

	stack "stack.example.com/api/v1"
)

func main() {
	s := stack.New()

	if s.IsEmpty() {
		fmt.Println("Stack is empty currently")
	}

	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(9)
	s.Print()
	fmt.Println("Pop the first element: ", s.Pop())
	fmt.Println("Top element in stack currently: ", s.Peek())
	fmt.Println("Is stack empty: ", s.IsEmpty())
	s.Print()
}
