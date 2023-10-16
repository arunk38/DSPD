package main

/*
 * Evaluate the value of the postfix expression.
 * Example:
 *	Input: 12 3 * 
 *	Output: 36 [12 * 3]
 *
 *	Input: 2 3 10 * + 9 -
 *	Output: 23 [(3*10) + 2 - 9]
 *
 * Approach:
 *	- Create a stack to store operands.
 *	- Scan the given expression and do following for every scanned element.
 *		- If the element is a number, push it into the stack
 *		- If the element is a operator, pop operands for the operator from stack. Evaluate the operator and push the result back to the stack
 *	- When the expression is ended, the number in the stack is the final answer
 */

import (
	"strings"
	"fmt"
	"strconv"
)

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value int
		prev *node
	}	
)

// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}

// Returns true if stack is empty, else false.
func (this *Stack) isEmpty() bool {
	if this.length == 0 {
		return true
	}
	return false
}

// View the top item on the stack
func (this *Stack) Peek() int {
	if this.length == 0 {
		return -1
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() int {
	if this.length == 0 {
		return -1
	}
	
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value int) {
	n := &node{value, this.top}
	this.top = n
	this.length++
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

func evalPostfix(expression string) {
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	expr := strings.Split(expression, " ")

	s := New()
	fmt.Print("Expression: ", expr)
	for _, i := range expr {
		if x, err := strconv.Atoi(i); err == nil {
			s.Push(x)
		} else {
			val1 := s.Pop()
			val2 := s.Pop()
			s.Push(ops[i](val2, val1))
		}
	}
	fmt.Println("\t\t\tvalue: ", s.Pop())
}

func main() {
	evalPostfix("12 3 *")
	evalPostfix("2 3 10 * + 9 -")
}