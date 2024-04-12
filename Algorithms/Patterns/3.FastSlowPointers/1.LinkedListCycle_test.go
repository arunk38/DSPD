package fastslowpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hasCycle(head *List, insertCycle bool) bool {

	if insertCycle {
		head.head.next.next.next = head.head.next.next
	}
	slow := head.head
	fast := head.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if slow == fast {
			return true
		}
	}

	return false
}

func TestLinkedListCycle(t *testing.T) {

	type args struct {
		arg1 []int
		arg2 bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1, 2, 4, 5, 6]",
			args: args{[]int{1, 2, 4, 5, 6}, false},
			want: false,
		},
		{
			name: "[1, 2, 4, 5, 6]",
			args: args{[]int{1, 2, 4, 5, 6}, true},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := List{}
			head.insertBulk(tt.args.arg1)
			got := hasCycle(&head, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
