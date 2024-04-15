package mergeintervals_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func insert(arr [][]int, newInterval []int) [][]int {
	if len(arr) < 2 {
		return arr
	}

	mergedIntervals := [][]int{}

	i := 0
	// skip all intervals that come before new interval
	for i < len(arr) && arr[i][1] < newInterval[0] {
		mergedIntervals = append(mergedIntervals, arr[i])
		i++
	}

	// merge all intervals that overlap with new interval
	for i < len(arr) && arr[i][0] <= newInterval[1] {
		newInterval[0] = min(arr[i][0], newInterval[0])
		newInterval[1] = max(arr[i][1], newInterval[1])
		i++
	}

	// insert the interval
	mergedIntervals = append(mergedIntervals, newInterval)

	// intert the remaining intervals
	for i < len(arr) {
		mergedIntervals = append(mergedIntervals, arr[i])
		i++
	}

	return mergedIntervals
}

func TestInsertInterval(t *testing.T) {

	type args struct {
		arg1 [][]int
		arg2 []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,6]",
			args: args{[][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 6}},
			want: [][]int{{1, 3}, {4, 7}, {8, 12}},
		},
		{
			name: "Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,10]",
			args: args{[][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 10}},
			want: [][]int{{1, 3}, {4, 12}},
		},
		{
			name: "Intervals=[[2,3],[5,7]], New Interval=[1,4]",
			args: args{[][]int{{2, 3}, {5, 7}}, []int{1, 4}},
			want: [][]int{{1, 4}, {5, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
