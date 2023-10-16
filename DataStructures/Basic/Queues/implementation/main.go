package main

import (
	"fmt"

	queue "queue.example.com/implementation/api"
)

func main() {
	q := queue.New()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)

	q.Print()
	fmt.Println("Front element in queue: ", q.FrontEle())
	fmt.Println("Front element in queue: ", q.RearEle())
	fmt.Println("Dequeue the first element: ", q.Dequeue())
	q.Print()
	fmt.Println("Is queue empty? ", q.IsEmpty())
}
