package main

/*
 * Given an array, print the Next Greater Element (NGE) for every element.
 * The Next greater Element for an element x is the first greater element on the right side of x in array.
 * Elements for which no greater element exist, consider next greater element as -1.
 *
 * Example:
 *	input: [4, 5, 2, 25]
 *	output:
 *  	Element ------ NGE
 *		   4	----->  5
 *		   5	----->  25
 *		   2	----->  25
 *		   25	----->  -1
 *
 *
 * Appraoch:
 *	Push all elemets to stack if their NGE is not found, loop over every element by popping and checking if loop element is NGE else push back
 *	Steps:
 *		- Push the first element to stack.
 *		- Pick rest of the elements one by one and follow the following steps in loop.
 *			- Mark the current element as current.
 *			- If stack is not empty, then pop an element from stack as popped
 *				- If popped is smaller than current, then current is NGE for popped element.
 *				- keep popping while popped elements are smaller than current and stack is not empty.
 *				- If popped is greater than next i.e., current, then push the popped back to stack.
 *			- Finally, push current to stack so that we can find next greater for it.
 *		- After iterating over the loop, the remaining elements in stack do not have the next greater so -1 for them.
 */

import (
	"fmt"

	stack "stack.example.com/api/v1"
)

func nge(input []int) {

	s := stack.New()
	s.Push(input[0])

	fmt.Println("\nInput array: ", input)
	fmt.Printf("%6s  ----%6s\n", "Eelement", "NGE")
	for _, ele := range input[1:] {
		curr := ele

		// If stack is not empty, then pop an element from stack as popped
		if !s.IsEmpty() {
			popped := s.Pop().(int)

			/*
			 * If popped is smaller than current, then current is NGE for popped element,
			 * keep popping while popped elements are smaller than current and stack is not empty
			 */
			for popped < curr {
				fmt.Printf("%8d  --->%6d\n", popped, curr)
				if s.IsEmpty() {
					break
				}
				popped = s.Pop().(int)
			}

			// If popped is greater than next, then push the popped back to stack
			if popped > curr {
				s.Push(popped)
			}
		}
		// Push current to stack so that we can find next greater for it.
		s.Push(curr)
	}

	/*
	 * After iterating over the loop, the remaining elements in stack do not have the next greater so -1 for them
	 */
	for !s.IsEmpty() {
		fmt.Printf("%8d  --->%6d\n", s.Pop(), -1)
	}
}

func main() {
	nge([]int{4, 5, 2, 25})
	nge([]int{13, 7, 6, 12})
	nge([]int{11, 13, 21, 3})
	nge([]int{11, 9, 7, 3})
}
