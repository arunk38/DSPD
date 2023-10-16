package main

/*
 *
 * Given a linked list, write a function to reverse every k nodes.
 * Example:
 *	Inputs: 1->2->3->4->5->6->7->8->NULL and k = 3
 *	Output: 3->2->1->6->5->4->8->7->NULL
 *
 * Approach:
 *	- Take prev as None, current_node as start node and an empty stack to store at most k items.
 *	- Till the current_node is not None:
 *		- Take at most k elements and insertAtStart and keep moving the current_node.
 *		- modify list head to current group head for first group
 *		- for later groups, loop til end of previous group and attach the next pointer to current group
 *
 *	Complexity:
 *		- Time: O(n)
 *		- Aux. Space: O(n)
 */

import (
	"fmt"
)

type Node struct {
	info int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) reverseInGroupSize(k int) {
	current_node := l.head
	first_group := true

	for current_node != nil {
		count := 0
		curr_group := List{}

		for (count < k) && current_node != nil {
			curr_group.insertAtStart(current_node.info)
			count++
			current_node = current_node.next
		}

		if first_group {
			l.head = result.head
			first_group = false
		} else {
			head := l.head
			for head.next != nil {
				head = head.next
			}
			head.next = curr_group.head
		}
	}
}

func (l *List) insertAtStart(d int) {
	list := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		list.next = l.head
		l.head = list
	}
}

func (l *List) insertAtEnd(d int) {
	list := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}

func printList(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
	fmt.Printf("-> nil\n")
}

func main() {
	num1 := List{}
	num1.insertAtEnd(1)
	num1.insertAtEnd(2)
	num1.insertAtEnd(3)
	num1.insertAtEnd(4)
	num1.insertAtEnd(5)
	num1.insertAtEnd(6)
	num1.insertAtEnd(7)
	num1.insertAtEnd(8)
	fmt.Println("Linked List at Start:")
	printList(&num1)

	fmt.Println("\nAfter reverse in group of 3:")
	num1.reverseInGroupSize(3)
	printList(&num1)

	num2 := List{}
	num2.insertAtEnd(1)
	num2.insertAtEnd(2)
	num2.insertAtEnd(3)
	num2.insertAtEnd(4)
	num2.insertAtEnd(5)
	num2.insertAtEnd(6)
	num2.insertAtEnd(7)
	num2.insertAtEnd(8)
	num2.insertAtEnd(9)
	fmt.Println("\nLinked List at Start:")
	printList(&num2)

	fmt.Println("\nAfter reverse in group of 5:")
	num2.reverseInGroupSize(5)
	printList(&num2)
}
