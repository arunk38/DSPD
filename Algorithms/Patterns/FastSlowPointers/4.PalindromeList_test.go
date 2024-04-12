package fastslowpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(ll **Node) {
	var prev, next, curr *Node
	curr = *ll

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	*ll = prev
}

func isPalindrome(l *List) bool {

	// find the middle node of the list
	middleNode := getMiddleNode(l)

	// reverse the second half
	reverse(&middleNode)

	p1 := l.head
	p2 := middleNode

	// compare both halves
	for p1 != nil && p2 != nil {
		if p1.info != p2.info {
			break // not a palindrome
		}

		p1 = p1.next
		p2 = p2.next
	}
	reverse(&middleNode) // revert the second half to its original state
	if p1 == nil || p2 == nil {
		return true
	}
	return false
}

func TestPalindromeList(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[2, 4, 6, 4, 2]",
			args: args{[]int{2, 4, 6, 4, 2}},
			want: true,
		},
		{
			name: "[2, 2, 4, 6, 4, 2]",
			args: args{[]int{2, 2, 4, 6, 4, 2}},
			want: false,
		},
		{
			name: "[1, 2, 3, 4, 5, 6, 7, 8, 9]",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := List{}
			l1.insertBulk(tt.args.arg1)
			got := isPalindrome(&l1)
			assert.Equal(t, tt.want, got)
		})
	}
}
