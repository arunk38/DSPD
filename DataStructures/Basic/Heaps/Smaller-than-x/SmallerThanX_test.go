package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Testing for smaller than function
 */

func TestSmallerThan(t *testing.T) {
	tests := []struct { // test case definition
		name string
		args bool
		got  any
		want any
	}{ // test cases
		{ // test case
			name: "Get smaller than 3 from {6, 5, 4, 8, 9, 10, 13, 12, 11, 7}",
			args: false,
			got:  smallerThan([]int{1, 23, 12, 9, 30, 2, 50}, 3),
			want: []int{1, 2},
		},
		{ // test case
			name: "Get smaller than 10 from {16, 23, 12, 19, 30, 42, 50}",
			args: false,
			got:  smallerThan([]int{16, 23, 12, 19, 30, 42, 50}, 10),
			want: []int{},
		},
	}

	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) { // function call
			assert.Equal(t, tt.want, tt.got) // test case assertions
		})
	}
}
