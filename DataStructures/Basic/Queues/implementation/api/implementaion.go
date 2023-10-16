package api

import "fmt"

/*
 * Implementation go queues with basic functionalities
 */

type (
	Queue struct {
		front  *node
		length int
	}
	node struct {
		value int
		next  *node
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil, 0}
}

// Returns true if queue is empty, else false.
func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

// Adds an item to the front of queue.
func (q *Queue) Enqueue(value int) {
	n := &node{value, nil}

	// if queue is empty, add current to front and return
	if q.front == nil {
		q.front = n
		q.length++
		fmt.Print("enqueueing ", value)
		q.Print()
		return
	}

	// loop to end of queue and add
	temp := q.front
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = n
	fmt.Print("enqueueing ", value)
	q.Print()
	q.length++
}

// Remove an item to the end of queue.
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	ele := q.front.value
	q.front = q.front.next

	return ele
}

func (q *Queue) FrontEle() int {
	if q.IsEmpty() {
		return -1
	}

	return q.front.value
}

func (q *Queue) RearEle() int {
	if q.IsEmpty() {
		return -1
	}

	temp := q.front
	for temp.next != nil {
		temp = temp.next
	}

	return temp.value
}

func (q *Queue) Print() {
	fmt.Print("   Queue curren state ---> [")
	n := q.front

	for n != nil {
		fmt.Print(" ", n.value)
		n = n.next
	}
	fmt.Print(" ]\n")
}
