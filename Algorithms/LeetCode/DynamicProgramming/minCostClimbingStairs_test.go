package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 *
 */

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	} else if n == 1 {
		return cost[0]
	}

	// create an array to store minimum cost to reach eash step
	dp := make([]int, n)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		// the minimum cost to reach the current step is the cost of
		// current step plus the minimum cost of reaching the previous 2 steps
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	// minimum cost to reach the top is the minimum of last 2 steps
	return min(dp[n-1], dp[n-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinCostClimbingStairs(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[10,15,20]",
			args: args{[]int{10, 15, 20}},
			want: 15,
		},
		{
			name: "[1,100,1,1,1,100,1,1,100,1]",
			args: args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostClimbingStairs(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
