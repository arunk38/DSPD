package main

import (
	queue "queue.example.com/implementation/api"
)

func main() {
	q := queue.New()
	q.Enqueue(10)
	q.Enqueue(20)

	q.Print()
}
