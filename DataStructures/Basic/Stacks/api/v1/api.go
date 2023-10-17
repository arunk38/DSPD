package api

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
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

// View the top item on the stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

// Print the contents of current stack
func (s *Stack) Print() {
	fmt.Print("[")
	n := s.top
	for n != nil {
		fmt.Print(" ", n.value)
		n = n.prev
	}
	fmt.Print(" ]\n")
}
