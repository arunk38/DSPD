package make_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given a sorted array, create a new array containing squares of
 * all the number of the input array in the sorted order.
 *
 *	Example 1:
 *	  Input: [-2, -1, 0, 2, 3]
 *	  Output: [0, 1, 4, 4, 9]
 *
 *	Example 2:
 *	  Input: [-3, -1, 0, 1, 2]
 *	  Output: [0, 1, 1, 4, 9]
 */

func makeSquares(input []int) []int {

	n := len(input)
	squareArray := []int{}
	left := 0
	right := n - 1

	for left <= right {
		leftSquare := input[left] * input[left]
		rigthSquare := input[right] * input[right]

		if leftSquare > rigthSquare {
			squareArray = append([]int{leftSquare}, squareArray...)
			left++
		} else {
			squareArray = append([]int{rigthSquare}, squareArray...)
			right--
		}
	}

	return squareArray
}

func TestMakeSqaures(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "make squares [-2, -1, 0, 2, 3]",
			args: args{[]int{-2, -1, 0, 2, 3}},
			want: []int{0, 1, 4, 4, 9},
		},
		{
			name: "make squares [-3, -1, 0, 1, 2]",
			args: args{[]int{-3, -1, 0, 1, 2}},
			want: []int{0, 1, 1, 4, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeSquares(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
