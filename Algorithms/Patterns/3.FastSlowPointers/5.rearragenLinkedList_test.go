package fastslowpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func rearrangeList(l *List) []int {

	// find the middle node of the list
	middleNode := getMiddleNode(l)

	// reverse the second half
	reverse(&middleNode)

	p1 := l.head
	p2 := middleNode

	for p1 != nil && p2 != nil {
		temp := p1.next
		p1.next = p2
		p1 = temp

		temp = p2.next
		p2.next = p1
		p2 = temp
	}

	// set the next of last node to 'nil'
	if p1 != nil {
		p1.next = nil
	}

	return l.convertToArray()
}

func TestRearrangeLinkedList(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[2, 4, 6, 8, 10, 12]",
			args: args{[]int{12, 10, 8, 6, 4, 2}},
			want: []int{2, 12, 4, 10, 6, 8},
		},
		{
			name: "[2, 4, 6, 8, 10]",
			args: args{[]int{10, 8, 6, 4, 2}},
			want: []int{2, 10, 4, 8, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := List{}
			l1.insertBulk(tt.args.arg1)
			got := rearrangeList(&l1)
			assert.Equal(t, tt.want, got)
		})
	}
}
