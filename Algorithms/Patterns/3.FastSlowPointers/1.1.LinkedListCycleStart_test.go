package fastslowpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func calcCycleLength(slow *Node) int {
	curr := slow
	cycleLength := 0

	for curr != slow {
		curr = curr.next
		cycleLength++
	}

	return cycleLength
}

func findStart(head *Node, len int) *Node {
	p1, p2 := head, head

	for len > 0 {
		p2 = p2.next
		len--
	}

	for p1 != p2 {
		p1 = p1.next
		p2 = p2.next
	}

	return p1
}

func cycleStart(head *List, insertCycle bool) *Node {

	if insertCycle {
		head.head.next.next.next = head.head.next.next
	}
	slow := head.head
	fast := head.head
	cycleLength := -1

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if slow == fast {
			cycleLength = calcCycleLength(slow)
			break
		}
	}

	if cycleLength == -1 {
		return nil
	}

	return findStart(head.head, cycleLength)
}

func TestLinkedListCycleStart(t *testing.T) {

	type args struct {
		arg1 []int
		arg2 bool
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1, 2, 4, 5, 6]",
			args: args{[]int{1, 2, 4, 5, 6}, false},
			want: -1,
		},
		{
			name: "[1, 2, 4, 5, 6]",
			args: args{[]int{1, 2, 4, 5, 6}, true},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := List{}
			head.insertBulk(tt.args.arg1)
			val := 0
			got := cycleStart(&head, tt.args.arg2)
			if got == nil {
				val = -1
			} else {
				val = got.info
			}
			assert.Equal(t, tt.want, val)
		})
	}
}
