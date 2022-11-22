package main

import (
	"math/rand"
)

func createRandomArray(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = rand.Int()
	}
	return res
}

func mergeSort(arr []int) []int {
	// check if len of arr > 1 and find mid, left and right arrays
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := arr[:mid]
		right := arr[mid:]

		// recurse on left and right
		left = mergeSort(left)
		right = mergeSort(right)

		// i for left and j for right
		i := 0
		j := 0
		// k for the main arr
		k := 0
		res := make([]int, len(arr))

		// replace arr valurs as per sorted values
		for i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				res[k] = left[i]
				i++
			} else {
				res[k] = right[j]
				j++
			}
			k++
		}

		// if there are any remaining values (for example, in case right and left arrays have values left)
		for i < len(left) {
			res[k] = left[i]
			i++
			k++
		}
		for j < len(right) {
			res[k] = right[j]
			j++
			k++
		}
		return res
	}
	return arr
}

func main() {
	n := 10000000
	arr := createRandomArray(n)
	_ = mergeSort(arr)
}
