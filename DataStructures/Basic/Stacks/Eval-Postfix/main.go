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
	"fmt"
	"strconv"
	"strings"

	stack "stack.example.com/api/v1"
)

func evalPostfix(expression string) {
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	expr := strings.Split(expression, " ")

	s := stack.New()
	fmt.Print("Expression: ", expr)
	for _, i := range expr {
		if x, err := strconv.Atoi(i); err == nil {
			s.Push(x)
		} else {
			val1 := s.Pop().(int)
			val2 := s.Pop().(int)
			s.Push(ops[i](val2, val1))
		}
	}
	fmt.Println("\t\t\tvalue: ", s.Pop())
}

func main() {
	evalPostfix("12 3 *")
	evalPostfix("2 3 10 * + 9 -")
}
