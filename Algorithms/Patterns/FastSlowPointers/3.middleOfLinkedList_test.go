package fastslowpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getMiddleNode(l *List) *Node {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func getMiddle(l *List) int {
	return getMiddleNode(l).info
}

func TestGetLinkedListMiddle(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[5, 4, 3, 2, 1]",
			args: args{[]int{5, 4, 3, 2, 1}},
			want: 3,
		},
		{
			name: "[6, 5, 4, 3, 2, 1]",
			args: args{[]int{6, 5, 4, 3, 2, 1}},
			want: 4,
		},

		{
			name: "[7, 6, 5, 4, 3, 2, 1]",
			args: args{[]int{7, 6, 5, 4, 3, 2, 1}},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := List{}
			l1.insertBulk(tt.args.arg1)
			got := getMiddle(&l1)
			assert.Equal(t, tt.want, got)
		})
	}
}
