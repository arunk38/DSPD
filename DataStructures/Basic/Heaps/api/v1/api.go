package api

import "fmt"

/*
 * Functionalities of heap
 */

type (
	Heap[T any] struct {
		data     []T
		comp     func(a, b T) bool
		heapSize int
	}
)

// Initialize a new heap with its comaparision function
func NewHeap[T any](comp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data:     make([]T, 0),
		comp:     comp,
		heapSize: 0,
	}
}

// Pop return the top value of the heap and rearrange the heap
// If the heap is empty, return zero value and false
func (h *Heap[T]) Pop() *Heap[T] {
	if h.GetHeapSize() == 0 {
		return nil
	}
	h.swap(0, h.GetHeapSize()-1)
	h.data = h.data[:h.GetHeapSize()-1]
	h.heapifyDown(0)
	h.heapSize--

	return h
}

// heapifyDown moves the value at index i down to its correct position in the heap
// Assume the left and right branch are already in the correct position in the heap except for i
// assumes that the binary trees rooted at LEFT.i/ and RIGHT.i/ are max- heaps, but that data at i
// might fail comparition thus violating the heap property
func (h *Heap[T]) heapifyDown(idx int) {

	leftChild := leftChildIdx(idx)
	rightChild := rightChildIdx(idx)

	largest := idx

	// See if left child of root exists and passes comparition
	if leftChild < len(h.data) && h.comp(h.data[idx], h.data[leftChild]) {
		largest = leftChild
	}

	// See if right child of root exists and passes comparition
	if rightChild < len(h.data) && h.comp(h.data[largest], h.data[rightChild]) {
		largest = rightChild
	}

	// swap if current index is modified
	if largest != idx {
		h.swap(idx, largest)
		h.heapifyDown(largest)
	}
}

// heapifyUp moves the value at index i up to its correct position in the heap
// Assume other values are already in the correct position in the heap except for i
func (h *Heap[T]) heapifyUp(i int) *Heap[T] {

	for h.comp(h.data[parentIdx(i)], h.data[i]) {
		h.swap(i, parentIdx(i))
		i = parentIdx(i)
	}

	return h
}

// Push an value into the already sorted heap
func (h *Heap[T]) Push(v T) *Heap[T] {
	h.data = append(h.data, v)
	h.heapSize++
	return h.heapifyUp((len(h.data) - 1))
}

// heapify an array, and return new heap obj
func (h *Heap[T]) BuildHeap(input []T) *Heap[T] {

	newHeap := NewHeap(h.comp)
	for _, v := range input {
		newHeap = newHeap.Push(v)
	}
	return newHeap
}

// deletes data to heap
func (h *Heap[T]) DeleteData() *Heap[T] {
	h.data = h.data[1:]
	h.heapSize--

	return h
}

// Prints the heap is array format
func (h *Heap[T]) PrintHeap() {
	fmt.Println("Current heap contents: ", h.data)
}

// sorts the given heap according the comparition function
// and returned sorted list
func (h *Heap[T]) HeapSort() []T {
	var sortedData []T

	tempData := h.data
	n := h.GetHeapSize()

	for len(sortedData) < n {
		sortedData = append(sortedData, h.data[0])
		h = h.BuildHeap(h.data[1:])
	}
	h.data = tempData

	return sortedData
}

func isHeap[T any](input []T, comp func(a, b T) bool, i int, n int) bool {

	// If (2 * i) + 1 >= n, then leaf node, so return true
	if i >= int((n-1)/2) {
		return true
	}

	// comparision should be done in reverse to what is done in heapify
	if comp(input[leftChildIdx(i)], input[i]) && comp(input[rightChildIdx(i)], input[i]) &&
		isHeap(input, comp, leftChildIdx(i), n) && isHeap(input, comp, rightChildIdx(i), n) {
		return true
	}
	return false
}

// check if the gven array input satisfies heap
func IsHeap[T any](input []T, comp func(a, b T) bool) bool {
	return isHeap(input, comp, 0, len(input))
}

// swap swaps the values at indices i and j
func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// swap swaps the values at indices i and j
func (h *Heap[T]) Peek() T {
	return h.data[0]
}

// return the left child index of the current heap element
func leftChildIdx(i int) int {
	return 2*i + 1
}

// return the right child index of the current heap element
func rightChildIdx(i int) int {
	return 2*i + 2
}

// return the parent element index for current element
func parentIdx(i int) int {
	return (i - 1) / 2
}

func (h *Heap[T]) GetHeapSize() int {
	return h.heapSize
}

func (h *Heap[T]) GetHeapData() []T {
	return h.data
}

// TODO: Delete
func Add(a, b int) int {
	return a + b
}
