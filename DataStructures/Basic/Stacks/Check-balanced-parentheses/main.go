package main

/*
 * Given an expression string, write a program to examine whether the pairs and the orders of "{", "}", "(", ")", "[", "]" are correct in expression.
 * Example:
 *	Input: "[()]{}{[()()]()}"
 *	Output: Balanced
 *
 * 	Input: "[(])"
 *	Output: Not Balanced
 *
 * Approach:
 *	- Initialize an empty stack
 *	- traverse the whole expression
 *		- if the bracket is opening bracket - '(', '{', '[' push to stack
 *		- if it is closing bracket  - ')', '}', ']', pop from stack and see if it matches its opening else parenthesis are not balanced.
 *	- After complete traversal, if there is some starting bracket left in stack then “not balanced”
 *
 *	Complexity:
 *		- Time: O(n)
 *		- Aux. Space: O(n)
 */

import (
	"fmt"
)

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value rune
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
func (this *Stack) Peek() rune {
	if this.length == 0 {
		return -1
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() rune {
	if this.length == 0 {
		return -1
	}
	
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value rune) {
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

func checkBalanced(expr string) {
	s := New()

	fmt.Println("\nInput parenthesis: ", expr)

	for _, char := range expr {
		// If current is opening bracket push closing bracket to stack
		if char == '(' {
			s.Push(')')
		} else if char == '{' {
			s.Push('}')
		} else if char == '[' {
			s.Push(']')
		} else if (s.isEmpty() == true) || (s.Pop() != char){
			/*
			 * If current is not opening bracket, and stack is empty or the bracket doesnt match the top of stack -> Unbalanced
			 */

			fmt.Println("Not balanced")
			return
		}
	}

	if s.isEmpty() == true {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not balanced")
	}
}

func main() {
	checkBalanced("[()]{}{[()()]()}")
	checkBalanced("[(])")
	checkBalanced("[](){")
}