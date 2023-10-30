package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/take-gifts-from-the-richest-pile/?envType=daily-question&envId=2023-10-28
// easy

func main() {
	gifts := []int{2, 25}
	k := 4
	pickGifts(gifts, k)
	fmt.Println(gifts)
}

func pickGifts(gifts []int, k int) int64 {
	// Write your code here
	maximumHeap(gifts)
	// take the square root of max heap every time in k times
	for i := 0; i < k; i++ {
		gifts[0] = int(math.Sqrt(float64(gifts[0])))
		heapify(gifts, 0, func(a, b int) bool {
			return a < b
		})
	}
	// sum the rest of the heap
	var sum int64
	for _, v := range gifts {
		sum += int64(v)
	}
	return sum
}

// Construct a maximum heap in generic way
func maximumHeap(arr []int) {
	// Start from bottom-most and rightmost
	// internal mode and heapify all internal
	// modes in bottom up way
	myLessFunc := func(a, b int) bool {
		return a < b
	}
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, myLessFunc)
	}
}

// To heapify a subtree rooted with node i which is
// an index in arr[].Nn is size of heap
func heapify[T any](arr []T, i int, lessFunc func(a, b T) bool) {
	n := len(arr)
	largest := i // Initialize largest as root
	l := 2*i + 1 // left = 2*i + 1
	r := 2*i + 2 // right = 2*i + 2

	// If left child is larger than root
	if l < n && lessFunc(arr[largest], arr[l]) {
		largest = l
	}

	// If right child is larger than largest so far
	if r < n && lessFunc(arr[largest], arr[r]) {
		largest = r
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected sub-tree
		heapify(arr, largest, lessFunc)
	}
}
