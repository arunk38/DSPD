package main

/*
 * Given a stack data structure with push and pop operations.
 * The task is to implement a queue using instances of stack data structure and operations on them.
 *
 * Approach:
 *	- Use 2 stacks to implement
 *	- enqueue: O(n)
 *		- While dataStack is not empty, push everything from dataStack to tempStack.
 *			- add at bottom of the dataStack
 *		- Push x to dataStack
 *		- Push everything back to stack1.
 *	- dequeue: O(1)
 *		- If dataStack is empty then error.
 *		- Pop an item from dataStack and return it.
 *	- front element:
 *		- top of the stack element is first element, peek it
 *	- rear element:
 *		- move all the elements from dataStack to tempStack
 *		- top of the tempStack will be last entry to queue, peek it
 *		- move back all elements to dataStack
 */

import (
	"fmt"

	stack "stack.example.com/api/v1"
)

type (
	Queue struct {
		dataStack stack.Stack
		tempStack stack.Stack
	}
)

func New() *Queue {
	return &Queue{*stack.New(), *stack.New()}

}

// Returns true if queue is empty, else false.
func (q *Queue) isEmpty() bool {
	return q.dataStack.IsEmpty()
}

// Adds an item to the front of queue.
func (q *Queue) enqueue(value interface{}) {
	for !q.dataStack.IsEmpty() {
		q.tempStack.Push(q.dataStack.Pop())
	}

	q.dataStack.Push(value)

	for !q.tempStack.IsEmpty() {
		q.dataStack.Push(q.tempStack.Pop())
	}
	fmt.Print("enqueueing ", value, " Current queue --> ")
	q.print()

}

// Remove an item to the end of queue.
func (q *Queue) dequeue() interface{} {
	if q.isEmpty() {
		fmt.Println("Empty queue")
		return nil
	}

	return q.dataStack.Pop()
}

func (q *Queue) frontEle() interface{} {
	if q.dataStack.IsEmpty() {
		return nil
	}

	return q.dataStack.Peek()
}

func (q *Queue) rearEle() interface{} {
	if q.dataStack.IsEmpty() {
		return nil
	}

	for !q.dataStack.IsEmpty() {
		q.tempStack.Push(q.dataStack.Pop())
	}

	val := q.tempStack.Peek()

	for !q.tempStack.IsEmpty() {
		q.dataStack.Push(q.tempStack.Pop())
	}

	return val
}

func (q *Queue) print() {
	q.dataStack.Print()
}

func main() {
	q := New()
	fmt.Println("Is queue empty? ", q.isEmpty())
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.enqueue(40)

	fmt.Println("Front element in queue: ", q.frontEle())
	q.print()
	fmt.Println("Rear element in queue: ", q.rearEle())
	fmt.Println("Dequeue the first element: ", q.dequeue())
	q.print()
	fmt.Println("Is queue empty? ", q.isEmpty())
	q.enqueue(120)
}
