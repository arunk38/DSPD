package main

/*
 * Given an infix expression conversation it to a postfix operation.
 * Example:
 *	input: a+b*(c-d)
 *	output: abcd-*+
 *
 *	input: a+b*(c^d-e)^(f+g*h)-i
 *	output: abcd^e-fgh*+^*+i-
 *
 * Appraoch:
 *	- Scan the infix expression from left to right.
 *	- If the scanned character is an operand, output it.
 *	- else,
 *		- If the precedence of the scanned operator is greater than the precedence of the
 *		  operator in the stack(or the stack is empty), push it.
 *		- Else, Pop the operator from the stack until the precedence of the scanned operator is
 *		  less-equal to the precedence of the operator residing on the top of the stack. Push the scanned operator to the stack.
 *	- If the scanned character is an ‘(‘, push it to the stack.
 *	- If the scanned character is an ‘)’, pop and output from the stack until an ‘(‘ is encountered.
 *	- Repeat above steps until infix expression is scanned.
 *	- Pop and output from the stack until it is not empty.
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

func infixToPostfix(expr string) {
	operator_precedence := map[rune]int{'+':1, '-':1, '*':2, '/':2, '%':2, '^':3}
	s := New()

	fmt.Print("Input expression: ", expr, "\t\t\t\tOutput: ")
	for _, char := range expr {
		// If scanned char is not an operand, then output it.
		if char >= 'a' && char <= 'z' {
			fmt.Printf("%c", char)
		} else if char == '(' {
			// Else if the scanned character is an ‘(‘, push it to the stack.
			s.Push(char)
		} else if char == ')' {
			// If the scanned character is an ‘)’, pop and output from the stack until an ‘(‘ is encountered.
			for s.isEmpty() == false && s.Peek() != '(' {
				fmt.Printf("%c", s.Pop())
			}
			s.Pop()
		} else {

			/*
			 * Pop the operator from the stack until operator_precedence[char] <= operator_precedence[stack.Peek]
			 * Push the scanned operator to the stack
			 */

			 for s.isEmpty() == false && s.Peek() != '(' && operator_precedence[char] <= operator_precedence[rune(s.Peek())] {
				fmt.Printf("%c", s.Pop())
			 }

			 /*
			  * If either stack is empty or precedence of scanned opeartor is greater
			  * than the top of stack, Push it to stack
			  */
			  s.Push(char)
		}
	}
	// Pop and output from the stack until it is not empty.
	for s.isEmpty() == false {
		fmt.Printf("%c", s.Pop())
	}
	fmt.Println()
}

func main() {
	infixToPostfix("a+b*(c-d)")
	infixToPostfix("a+b*(c^d-e)^(f+g*h)-i")
}