package golang

import "log"

/*
	1) Simple Solution
	In this we will simply just count number of zeroes in the list and then replace that amount of zeroes in the given arrays
	in the beginning.

	Rest of the elements in the array we can set as 1.

	Remember, we can also count 1s and then set, upto you.

	This solution will take 2n time, which is O(n) and space will be O(1).
*/

// NaiveSortBinaryArray ...
func NaiveSortBinaryArray(input []int) {
	z := 0

	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			z++
		}
	}

	pos := 0
	for z > 0 {
		input[pos] = 0
		z--
		pos++
	}

	for pos < len(input) {
		input[pos] = 1
		pos++
	}

	log.Printf("Final Array: %#v\n", input)
}

/*
	2) Partitioning Solution

	We can use quick-sort's partitioning solution to sort this array easily with given constraints.

	We can take 1 as pivot element and then go through array and swap all the elements less then 1 with 1.

	This will give us true O(n) time complexity as well as O(1) space complexity.
*/

// SortBinaryArray ...
func SortBinaryArray(input []int) {

	pivot := 1
	pivotIndex := 0

	for i := 0; i < len(input); i++ {
		if input[i] < pivot {
			Swap(input, i, pivotIndex)
			pivotIndex++
		}
	}

	log.Printf("Sorted: %#v\n", input)
}
