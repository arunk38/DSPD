package mergeintervals_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(arr [][]int) [][]int {
	if len(arr) < 2 {
		return arr
	}

	sort.Slice(arr, func(i, j int) bool {
		// sort based on start of interval
		return arr[i][0] < arr[j][0]
	})

	mergedInterval := [][]int{}

	start := arr[0][0]
	end := arr[0][1]

	for i := 1; i < len(arr); i++ {
		if arr[i][0] <= end {
			// overlapping intervals, adjust the end time
			end = max(arr[i][1], end)
		} else {
			// non-overlapping intervals, add the previous interval and reset start and end
			mergedInterval = append(mergedInterval, []int{start, end})
			start = arr[i][0]
			end = arr[i][1]
		}
	}

	// add the last interval
	mergedInterval = append(mergedInterval, []int{start, end})

	return mergedInterval
}

func TestOvelappingIntervals(t *testing.T) {

	type args struct {
		arg1 [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[[1,4], [2,5], [7,9]]",
			args: args{[][]int{{1, 4}, {2, 5}, {7, 9}}},
			want: [][]int{{1, 5}, {7, 9}},
		},
		{
			name: "[[6,7], [2,4], [5,9]]",
			args: args{[][]int{{6, 7}, {2, 4}, {5, 9}}},
			want: [][]int{{2, 4}, {5, 9}},
		},
		{
			name: "[[1,4], [2,6], [3,5]]",
			args: args{[][]int{{1, 4}, {2, 6}, {3, 5}}},
			want: [][]int{{1, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
