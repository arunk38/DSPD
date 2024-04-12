package twopointers

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given an array, find the length of the smallest subarray in it which when
 * sorted will sort the whole array.
 */

func minWindowSort(input []int) int {

	low := 0
	high := len(input) - 1

	// find the first number out of sorting from begining
	for low < high && input[low] <= input[low+1] {
		low++
	}

	// Array is already sorted
	if low == high {
		return 0
	}

	// find the first number out of sorting from end
	for high >= 0 && input[high] >= input[high-1] {
		high--
	}

	// find the max/min of the sub-array

	subArrayMax := math.MinInt64
	subArrayMin := math.MaxInt64

	for k := low; k < high; k++ {
		subArrayMax = int(math.Max(float64(subArrayMax), float64(input[k])))
		subArrayMin = int(math.Min(float64(subArrayMin), float64(input[k])))
	}

	// extend the sub-array to include any number which is bigger than the min of subarray
	for low > 0 && input[low-1] > subArrayMin {
		low--
	}

	// extend the sub-array to include any number which is smaller than the max of subarray
	for high < len(input)-1 && input[high+1] < subArrayMin {
		high++
	}

	return high - low + 1
}

func TestMinWindowSort(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1, 2, 5, 3, 7, 10, 9, 12]",
			args: args{[]int{1, 2, 5, 3, 7, 10, 9, 12}},
			want: 5,
		},
		{
			name: "[1, 3, 2, 0, -1, 7, 10]",
			args: args{[]int{1, 3, 2, 0, -1, 7, 10}},
			want: 5,
		},
		{
			name: "[1, 2, 3]",
			args: args{[]int{1, 2, 3}},
			want: 0,
		},
		{
			name: "[3, 2, 1]",
			args: args{[]int{3, 2, 1}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindowSort(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
