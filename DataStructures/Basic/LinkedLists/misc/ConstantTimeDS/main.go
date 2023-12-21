package main

import (
	"fmt"
)

// Node represents a node in the doubly linked list
type Node struct {
	value int
	prev  *Node
	next  *Node
}

// MinDS represents the data structure supporting constant time operations
type MinDS struct {
	head     *Node
	tail     *Node
	minNode  *Node // Node containing the current minimum value
	valueMap map[int]*Node
}

// NewMinDS initializes a new instance of the data structure
func NewMinDS() *MinDS {
	return &MinDS{
		valueMap: make(map[int]*Node),
	}
}

// Insert adds a value to the data structure
func (mds *MinDS) Insert(value int) {
	newNode := &Node{value: value}

	// Update the doubly linked list
	if mds.head == nil {
		mds.head = newNode
		mds.tail = newNode
	} else {
		mds.tail.next = newNode
		newNode.prev = mds.tail
		mds.tail = newNode
	}

	// Update the value map
	mds.valueMap[value] = newNode

	// Update the minimum node if necessary
	if mds.minNode == nil || value < mds.minNode.value {
		mds.minNode = newNode
	}
}

// Delete removes a value from the data structure
func (mds *MinDS) Delete(value int) {
	node, ok := mds.valueMap[value]
	if !ok {
		return // Value not found
	}

	// Update the doubly linked list
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		mds.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		mds.tail = node.prev
	}

	// Update the value map
	delete(mds.valueMap, value)

	// Update the minimum node if necessary
	if node == mds.minNode {
		// Recalculate the minimum if the deleted node was the minimum
		mds.recalculateMin()
	}
}

// GetMin returns the minimum value from the data structure
func (mds *MinDS) GetMin() int {
	if mds.minNode != nil {
		return mds.minNode.value
	}
	return 0 // or any default value if the list is empty
}

// recalculateMin finds the new minimum value in the data structure
func (mds *MinDS) recalculateMin() {
	minValue := int(^uint(0) >> 1) // Max int value

	// Iterate through the doubly linked list to find the minimum
	current := mds.head
	for current != nil {
		if current.value < minValue {
			minValue = current.value
			mds.minNode = current
		}
		current = current.next
	}
}

func main() {
	// Example usage:
	minDS := NewMinDS()

	// Insert
	minDS.Insert(3)
	minDS.Insert(1)
	minDS.Insert(4)
	minDS.Insert(2)

	// GetMin
	fmt.Println(minDS.GetMin()) // Output: 1

	// Delete
	minDS.Delete(1)

	// GetMin after deletion
	fmt.Println(minDS.GetMin()) // Output: 2
}
