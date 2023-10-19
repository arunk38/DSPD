package heap_test

import (
	"github.com/stretchr/testify/assert"
	heap "heaps.example.com/api/v1"

	"testing"
)

func TestHeapify(t *testing.T) {
	type Person struct {
		id   int
		name string
		age  int
	}
	tests := []struct { // test case definition
		name string
		args bool
		got  any
		want any
	}{ // test cases
		{ // test case
			name: "Max Heap",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}).GetHeapData(),
			want: []int{13, 12, 10, 11, 7, 4, 9, 5, 8, 6},
		},
		{ // test case
			name: "Get size of heap {6, 5, 4, 8, 9, 10, 13, 12, 11, 7}",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}).GetHeapSize(),
			want: 10,
		},
		{ // test case
			name: "Peek Max Heap",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}).Peek(),
			want: 13,
		},
		{ // test case
			name: "Sort in max Heap",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7, 1}).HeapSort(),
			want: []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 1},
		},
		{ // test case
			name: "Validate heap push",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7, 1}).Push(20).GetHeapData(),
			want: []int{20, 12, 13, 11, 7, 10, 9, 5, 8, 6, 1, 4},
		},
		{ // test case
			name: "Validate heap pop",
			args: true,
			got: heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7, 1}).Pop().GetHeapData(),
			want: []int{12, 11, 10, 8, 7, 4, 9, 5, 1, 6},
		},
		{ // test case
			name: "Min Heap",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a > b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}).GetHeapData(),
			want: []int{4, 6, 5, 8, 7, 10, 13, 12, 11, 9},
		},
		{ // test case
			name: "Sort in min Heap",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a > b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}).HeapSort(),
			want: []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		},
		{ // test case
			name: "Empty heap, size is zero",
			args: false,
			got:  heap.NewHeap(func(a, b int) bool { return a < b }).GetHeapSize(),
			want: 0,
		},
		{ // test case
			name: "Peek Min Heap",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a > b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7, 1}).Peek(),
			want: 1,
		},
		{ // test case
			name: "Check if array is Heap - min",
			args: false,
			got: heap.IsHeap(heap.NewHeap(func(a, b int) bool { return a > b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}).GetHeapData(), func(a, b int) bool { return a > b }),
			want: true,
		},
		{ // test case
			name: "Check if array is Heap - max",
			args: false,
			got: heap.IsHeap(heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}).GetHeapData(), func(a, b int) bool { return a < b }),
			want: true,
		},
		{ // test case
			name: "Check if array is Heap - invalid",
			args: false,
			got:  heap.IsHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7, 1}, func(a, b int) bool { return a > b }),
			want: false,
		},
		{ // test case
			name: "Get size of heap {5, 10}",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{5, 10}).GetHeapSize(),
			want: 2,
		},
		{ // test case
			name: "Delete from heap {5, 10}",
			args: false,
			got: heap.NewHeap(func(a, b int) bool { return a < b }).
				BuildHeap([]int{5, 10}).
				DeleteData().GetHeapSize(),
			want: 1,
		},
		{ // test case
			name: "Test heap with struct object",
			args: false,
			got: heap.NewHeap(func(a, b Person) bool { return a.age < b.age }).
				BuildHeap([]Person{
					{1, "Jerry", 20},
					{2, "Tome", 18},
					{3, "Nick", 19},
					{4, "Bob", 34},
					{5, "Tommy", 36},
				}).GetHeapData(),
			want: []Person{
				{5, "Tommy", 36},
				{4, "Bob", 34},
				{3, "Nick", 19},
				{2, "Tome", 18},
				{1, "Jerry", 20},
			},
		},
		{ // test case
			name: "Test heap with struct object",
			args: false,
			got: heap.NewHeap(func(a, b Person) bool { return a.age < b.age }).
				BuildHeap([]Person{
					{1, "Jerry", 20},
					{2, "Tome", 18},
					{3, "Nick", 19},
					{4, "Bob", 34},
					{5, "Tommy", 36},
				}).Pop().GetHeapData(),
			want: []Person{
				{4, "Bob", 34},
				{1, "Jerry", 20},
				{3, "Nick", 19},
				{2, "Tome", 18},
			},
		},
	}

	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) { // function call
			assert.Equal(t, tt.want, tt.got) // test case assertions
		})
	}
}
