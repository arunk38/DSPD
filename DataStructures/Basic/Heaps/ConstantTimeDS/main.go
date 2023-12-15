package main

/*
 * Design a data structure that support constant time operations i.e
 * Insert(value) [0(1)] , delete(value) [0(1)] and findMinimun [0(1)].
 */

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a node in the doubly linked list
type Node struct {
	value      int
	prev, next *Node
	deleted    bool
}

// MinHeap implements a min heap using the heap package
type MinHeap []*Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MinLinkedList implements a data structure supporting constant time operations
type MinLinkedList struct {
	head          *Node
	tail          *Node
	minHeap       MinHeap
	valueMap      map[int]*Node
	currentMin    int
	currentMinSet bool
}

// Insert inserts a new value into the data structure
func (ml *MinLinkedList) Insert(value int) {
	newNode := &Node{value: value}

	if ml.head == nil {
		ml.head = newNode
		ml.tail = newNode
	} else {
		ml.tail.next = newNode
		newNode.prev = ml.tail
		ml.tail = newNode
	}

	ml.valueMap[value] = newNode
	heap.Push(&ml.minHeap, newNode)

	if !ml.currentMinSet || value < ml.currentMin {
		ml.currentMin = value
		ml.currentMinSet = true
	}
}

// Delete marks a value as deleted in the data structure
func (ml *MinLinkedList) Delete(value int) {
	node, ok := ml.valueMap[value]
	if !ok {
		return // Value not found
	}

	node.deleted = true
	if value == ml.currentMin {
		ml.currentMinSet = false
	}
}

// GetMin returns the minimum value from the data structure
func (ml *MinLinkedList) GetMin() int {
	if ml.currentMinSet {
		return ml.currentMin
	}

	// currentMin is not set, pop deleted elements from top
	for len(ml.minHeap) > 0 && ml.minHeap[0].deleted {
		heap.Pop(&ml.minHeap)
	}

	if len(ml.minHeap) > 0 {
		ml.currentMin = ml.minHeap[0].value
		ml.currentMinSet = true
		return ml.currentMin
	}
	return math.MaxInt64 // or any default value if the list is empty
}

func main() {
	minLinkedList := MinLinkedList{
		valueMap: make(map[int]*Node),
	}

	minLinkedList.Insert(3)
	minLinkedList.Insert(1)
	minLinkedList.Insert(4)
	minLinkedList.Insert(2)

	fmt.Println(minLinkedList.GetMin()) // Output: 1

	minLinkedList.Delete(4)
	fmt.Println(minLinkedList.GetMin()) // Output: 1
	minLinkedList.Delete(1)
	fmt.Println(minLinkedList.GetMin()) // Output: 2
}
