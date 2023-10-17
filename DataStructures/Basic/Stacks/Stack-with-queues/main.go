package main

/*
 * Given a queue data structure with enqueue and dequeue operations.
 * The task is to implement a stack using instances of queue data structure and operations on them.
 *
 * Approach:
 *	- Use 2 queues to implement - dataQueue and tempQueue
 *	- Push: O(n)
 *		- enqueue to tempQueue
 *		- One by one dequeue everything from dataQueue and enqueue to tempQueue
 *		- Swap the dataQueue and tempQueue
 *	- Pop: O(1)
 *		- Dequeue an item from dataQueue and return it.
 *	- Peek:
 *		- rearEle of dataQueue will be top of stack
 */

import (
	"fmt"

	queue "queue.example.com/api/v1"
)

type (
	Stack struct {
		dataQueue queue.Queue
		tempQueue queue.Queue
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{*queue.New(), *queue.New()}
}

// Returns true if stack is empty, else false.
func (s *Stack) isEmpty() bool {
	return s.dataQueue.IsEmpty()
}

func (s *Stack) push(value interface{}) {
	// enqueue the data to dataQueue
	s.tempQueue.Enqueue(value)

	// move all elements from data to temp
	for !s.dataQueue.IsEmpty() {
		s.tempQueue.Enqueue(s.dataQueue.Dequeue())
	}

	// swap data and temp queues in stack
	s.dataQueue, s.tempQueue = s.tempQueue, s.dataQueue
}

// Pop the top item of the stack and return it
func (s *Stack) pop() interface{} {
	if s.dataQueue.IsEmpty() {
		return nil
	}

	return s.dataQueue.Dequeue()
}

// View the top item on the stack
func (s *Stack) peek() interface{} {
	return s.dataQueue.RearEle()
}

func (s *Stack) print() {
	s.dataQueue.Print()
}

func main() {

	s := New()

	if s.isEmpty() {
		fmt.Println("Stack is empty currently")
	}

	s.push(4)
	s.push(5)
	s.push(6)
	s.push(9)
	s.print()
	fmt.Println("Pop the first element: ", s.pop())
	fmt.Println("Top element in stack currently: ", s.peek())
	fmt.Println("Is stack empty: ", s.isEmpty())
	s.print()

}
