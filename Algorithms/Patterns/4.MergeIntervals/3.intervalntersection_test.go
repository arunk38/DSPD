package mergeintervals_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intersect(arr1 [][]int, arr2 [][]int) [][]int {

	result := [][]int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		// check if interval arr1[i] intersects with arr2[j]
		// check if one of the interval start time lies with in another
		if (arr1[i][0] >= arr2[j][0] && arr1[i][0] <= arr2[j][1]) ||
			(arr2[j][0] >= arr1[i][0] && arr2[j][0] <= arr1[i][1]) {
			// store the intersection part
			result = append(result, []int{max(arr1[i][0], arr2[j][0]), min(arr1[i][1], arr2[j][1])})
		}

		// move from interval which is finishing first
		if arr1[i][1] < arr2[j][1] {
			i++
		} else {
			j++
		}
	}
	return result
}

func TestIntervalIntersect(t *testing.T) {

	type args struct {
		arg1 [][]int
		arg2 [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "arr1=[[1, 3], [5, 6], [7, 9]], arr2=[[2, 3], [5, 7]]",
			args: args{[][]int{{1, 3}, {5, 6}, {7, 9}}, [][]int{{2, 3}, {5, 7}}},
			want: [][]int{{2, 3}, {5, 6}, {7, 7}},
		},
		{
			name: "arr1=[[1, 3], [5, 7], [9, 12]], arr2=[[5, 10]]",
			args: args{[][]int{{1, 3}, {5, 7}, {9, 12}}, [][]int{{5, 10}}},
			want: [][]int{{5, 7}, {9, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
