package main

/*
 *	Given a number n, write a function that generates and prints all binary numbers with decimal values from 1 to n.
 *	Example:
 *		Input: n = 2
 *		Output: 1, 10
 *
 *		Input: n = 5
 *		Output: 1, 10, 11, 100, 101
 *
 *	Approach:
 *		- Create an empty queue
 *		- Enqueue the first binary number “1” to queue.
 *		- Now run a loop for generating and printing n binary numbers
 *			- Dequeue and Print the front of queue.
 *			- append "0" and enqueue
 *			- append "1" and enqueue
 *
 */

import (
	"fmt"

	queue "queue.example.com/api/v1"
)

func generateNBinary(n int) {
	fmt.Println("\nN value", n)

	q := queue.New()

	q.Enqueue("1")

	for i := 0; i < n; i++ {
		s := q.Dequeue()
		fmt.Print(" ", s.(string))
		q.Enqueue(s.(string) + "0")
		q.Enqueue(s.(string) + "1")
	}
}

func main() {
	generateNBinary(3)
	generateNBinary(5)
	generateNBinary(8)
}
