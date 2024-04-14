package fastslowpointers

import "fmt"

type Node struct {
	info int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) convertToArray() []int {
	ret := []int{}
	temp := l.head

	for temp != nil {
		ret = append(ret, temp.info)
		temp = temp.next
	}
	return ret
}

func printList(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
	fmt.Printf("-> nil\n")
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

func (l *List) insertBulk(input []int) {

	for i, _ := range input {
		list := &Node{info: input[i], next: nil}
		if l.head == nil {
			l.head = list
		} else {
			list.next = l.head
			l.head = list
		}
	}
}
