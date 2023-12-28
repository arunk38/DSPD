package main

/*
 * A double booking happens when two events have some non-empty intersection
 * (i.e., some moment is common to both events.).
 */

import "fmt"

type Node struct {
	// start and end times of the event
	start, end  int
	left, right *Node
}

// MyCalendar represents a calender using a BST to manage events
type MyCalendar struct {
	root *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}

}

// inserts a new event into the BST rooted at the given node.
// returns true if the event was successfully booked, flase it there was a conflict
func insertNode(node *Node, start, end int) bool {
	// if the new event's end time is before the current node's start time,
	// insert in the left subtree.
	if end <= node.start {
		if node.left == nil {
			node.left = &Node{start: start, end: end}
			return true
		}

		// recursively insert into the left subtree
		return insertNode(node.left, start, end)
	} else if start >= node.end {
		// if the new event's end time is after the current node's end time,
		// insert in the right subtree.
		if node.right == nil {
			node.right = &Node{start: start, end: end}
			return true
		}

		// recursively insert into the left subtree
		return insertNode(node.right, start, end)
	} else {
		// Conflict: the new event overlaps with current node's interval
		return false
	}

}

// adds a new event to the calender.
// returns true if the event was successfully booked.
func (cal *MyCalendar) Book(start int, end int) bool {
	//if the calender is empty, create the first event as root.
	if cal.root == nil {
		cal.root = &Node{start: start, end: end}
		return true
	}

	// Otherwise, insert the new event into the BST
	return insertNode(cal.root, start, end)
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Book(10, 20))
	fmt.Println(obj.Book(15, 25))
	fmt.Println(obj.Book(20, 30))
}
