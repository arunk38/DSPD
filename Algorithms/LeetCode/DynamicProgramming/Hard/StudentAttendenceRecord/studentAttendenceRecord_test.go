package studentattendencerecord_test

/*
 * leetcode: https://leetcode.com/problems/student-attendance-record-ii/description/
 */

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkRecord(n int) int {
	mod := int(1e9) + 7

	// Initialize a 3D array to represent the state transitions
	// dp[i][j][k] represents the number of valid records for the ith day
	// with j absent days and ending with k consecutive late days
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
		}
	}

	// Initialize the base case for dp[0]
	// only 1 valid record - student not absent, not late on day 1
	dp[0][0][0] = 1

	// Iterate through the days and calculate state transitions
	for i := 1; i <= n; i++ {
		// Case 1: Include 'P' (Present)
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 2; k++ {
				// Calculate the number of valid records for the current day
				// by including 'P' and considering transitions from the previous day
				dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][k]) % mod
			}
		}

		// Case 2: Include 'A' (Absent)
		for k := 0; k <= 2; k++ {
			// Calculate the number of valid records for the current day
			// by including 'A' and considering transitions from the previous day
			dp[i][1][0] = (dp[i][1][0] + dp[i-1][0][k]) % mod
		}

		// Case 3: Include 'L' (Late)
		for j := 0; j <= 1; j++ {
			for k := 1; k <= 2; k++ {
				// Calculate the number of valid records for the current day
				// by including 'L' and considering transitions from the previous day
				dp[i][j][k] = (dp[i][j][k] + dp[i-1][j][k-1]) % mod
			}
		}
	}

	// Sum up the valid records for the last day across all possible states
	result := 0
	for j := 0; j <= 1; j++ {
		for k := 0; k <= 2; k++ {
			result = (result + dp[n][j][k]) % mod
		}
	}

	return result
}

func TestCheckRecord(t *testing.T) {

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
			want: 8,
		},
		{
			name: "n=1",
			args: args{1},
			want: 3,
		},
		{
			name: "n=10101",
			args: args{10101},
			want: 183236316,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkRecord(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
