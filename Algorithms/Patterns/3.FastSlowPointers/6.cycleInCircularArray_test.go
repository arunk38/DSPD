package fastslowpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findNextIndex(arr []int, isForward bool, currIndex int) int {

	direction := arr[currIndex] >= 0

	if isForward != direction {
		return -1 // change in direction, return -1
	}

	nextIndex := (currIndex + arr[currIndex]) % len(arr)
	if nextIndex < 0 {
		nextIndex += len(arr) // wrap around for negative numbers
	}

	if nextIndex == currIndex {
		return -1 // one element cycle
	}

	return nextIndex
}

func loopExists(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		isForward := arr[i] >= 0 // if we are moving foreard or not

		slow := i
		fast := i

		// simulationg do-while
		slow = findNextIndex(arr, isForward, slow)
		fast = findNextIndex(arr, isForward, fast)

		// move fast pointer twice
		if fast != -1 {
			fast = findNextIndex(arr, isForward, fast)
		}

		// if slow and fast becomes -1, means we cant find any cycle for this number
		for slow != -1 && fast != -1 && slow != fast {
			slow = findNextIndex(arr, isForward, slow)
			fast = findNextIndex(arr, isForward, fast)

			// move fast pointer twice
			if fast != -1 {
				fast = findNextIndex(arr, isForward, fast)
			}
		}

		if slow != -1 && slow == fast {
			return true
		}
	}

	return false
}

func TestCycleInCircularArray(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1, 2, -1, 2, 2]",
			args: args{[]int{1, 2, -1, 2, 2}},
			want: true,
		},
		{
			name: "[2, 2, -1, 2]",
			args: args{[]int{2, 2, -1, 2}},
			want: true,
		},
		{
			name: "[2, 1, -1, 2]",
			args: args{[]int{2, 1, -1, 2}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := loopExists(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
