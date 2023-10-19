package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Testing for klargest function
 */

func TestKLargest(t *testing.T) {
	tests := []struct { // test case definition
		name string
		args bool
		got  any
		want any
	}{ // test cases
		{ // test case
			name: "Get 3 largest, from {6, 5, 4, 8, 9, 10, 13, 12, 11, 7}",
			args: false,
			got:  kLargest([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}, 3),
			want: []int{13, 12, 11},
		},
		{ // test case
			name: "Get 0 largest",
			args: false,
			got:  kLargest([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}, 0),
			want: []int{},
		},
		{ // test case
			name: "Invalid value of k",
			args: false,
			got:  kLargest([]int{6, 5}, 4),
			want: []int{},
		},
	}

	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) { // function call
			assert.Equal(t, tt.want, tt.got) // test case assertions
		})
	}
}
