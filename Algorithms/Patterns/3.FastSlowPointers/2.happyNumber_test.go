package fastslowpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findSquareSum(num int) int {
	sum := 0
	var digit int

	for num > 0 {
		digit = num % 10
		sum += digit * digit
		num /= 10
	}

	return sum
}

func happyNumber(n int) bool {

	slow := findSquareSum(n)
	fast := findSquareSum(slow)

	for slow != fast {
		slow = findSquareSum(slow)
		fast = findSquareSum(findSquareSum(fast))
	}

	return slow == 1
}

func TestHapppyNumber(t *testing.T) {

	type args struct {
		arg1 int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "23",
			args: args{23},
			want: true,
		},
		{
			name: "12",
			args: args{12},
			want: false,
		},
		{
			name: "54",
			args: args{54},
			want: false,
		},
		{
			name: "32",
			args: args{32},
			want: true,
		},
		{
			name: "10",
			args: args{10},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := happyNumber(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
