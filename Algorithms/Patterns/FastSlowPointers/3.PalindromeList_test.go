package fastslowpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(l *List) bool {
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
			name: "[1, 2, 4, 5, 6]",
			args: args{[]int{2, 4, 6, 4, 2, 2}},
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
