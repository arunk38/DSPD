package climbstairs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 *
 */

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	curr, next := 2, 3

	for i := 3; i < n; i++ {
		/*
		 * Nth stair can be reached from (N-1)th stair by taking 1 step
		 * or from (N-2)th stair by taking 2 steps.
		 * Number of ways to reach Nth = ways (N-1)th + ways (N-2)th
		 * N = 4, can be reached from 3rd with 1 step, and 2nd with 2 steps
		 */
		curr, next = next, curr+next
	}

	return next
}

func TestClimbStairs(t *testing.T) {

	type args struct {
		arg1 int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=2",
			args: args{2},
			want: 2,
		},
		{
			name: "n=5",
			args: args{5},
			want: 8,
		},
		{
			name: "n=15",
			args: args{15},
			want: 987,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
