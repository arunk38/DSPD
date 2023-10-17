package main

/*
 *
 * Given two numbers represented by two lists, write a function that returns sum list.
 * The sum list is list representation of addition of two input numbers.
 * Eample:
 *	input:
 *		num1: 984 (4->8->9)
 *		num2 = 978 (8->7->9)
 *	output:
 *		sum = 1962 (2->6->9->1)
 *
 * Approach:
 *	- Take 2 pointers to track both the lists, take carry = 0 and also initialize result list.
 *	- Now while any of temp_1 or temp_2 is there keep going.
 *		- Get the sum: carry + first_list digit sum if exists else 0 + second_list digit sum if exists else 0.
 *		- If sum geater than 10 then get the digit and carry separated.
 *		- Push the digit to result list at end and update the temp_1 and temp_2 pointers if exists else None.
 *	- Push the digit to result list and update the temp_1 and temp_2 pointers if exists else None.
 *	- Lastly, set the head of the list as result list’s head.
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

func (l *List) insert(d int) {
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

func convToInt(head *Node, idx int) int {
	if head.next == nil {
		return idx * head.info
	}

	return idx*head.info + convToInt(head.next, idx*10)
}

func (l *List) add(num *List) {
	temp1 := l.head
	temp2 := num.head
	carry := 0
	result := List{}

	for temp1 != nil || temp2 != nil {
		sum := carry

		if temp1 != nil {
			sum += temp1.info
			temp1 = temp1.next
		}
		if temp2 != nil {
			sum += temp2.info
			temp2 = temp2.next
		}

		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		result.insertAtEnd(sum)
	}

	if carry != 0 {
		result.insertAtEnd(carry)
	}
	l.head = result.head
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
	num1.insert(6)
	num1.insert(4)
	num1.insert(9)
	num1.insert(5)
	num1.insert(7)
	printList(&num1)
	fmt.Println("Number 1:", convToInt(num1.head, 1))

	num2 := List{}
	num2.insert(4)
	num2.insert(8)
	printList(&num2)
	fmt.Println("Number 2:", convToInt(num2.head, 1))

	num1.add(&num2)
	printList(&num1)
	fmt.Println("Sum of both:", convToInt(num1.head, 1))
}
