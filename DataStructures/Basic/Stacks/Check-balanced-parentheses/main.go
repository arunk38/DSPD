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

	stack "stack.example.com/api/v1"
)

func checkBalanced(expr string) {
	s := stack.New()

	fmt.Println("\nInput parenthesis: ", expr)

	for _, char := range expr {
		// If current is opening bracket push closing bracket to stack
		if char == '(' {
			s.Push(')')
		} else if char == '{' {
			s.Push('}')
		} else if char == '[' {
			s.Push(']')
		} else if (s.IsEmpty()) || (s.Pop().(rune) != char) {
			/*
			 * If current is not opening bracket, and stack is empty or the bracket doesnt match the top of stack -> Unbalanced
			 */

			fmt.Println("Not balanced")
			return
		}
	}

	if s.IsEmpty() {
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
