package countingbits_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 *
 */

func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		// The number of set bits in `i` is equal to the number of set bits in i/2 plus the least significant bit of i
		result[i] = result[i/2] + i%2
	}

	return result
}

func TestCountBits(t *testing.T) {

	type args struct {
		arg1 int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "n=2",
			args: args{2},
			want: []int{0, 1, 1},
		},
		{
			name: "n=5",
			args: args{5},
			want: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "n=15",
			args: args{15},
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
