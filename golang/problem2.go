package golang

import "log"

/*
	1) Naive Solution

	For naive solution, the logic is simple.
	Start from element i, which will be the first element of sub-array, Now, go through other elements of the array
	and keep adding them to ith element and see if the sum is zero.

	If it is, then those are the elements which represents our output.

	For, more understanding, please check out the function below.

	This will give us time complexity of O(n^2) and space complexity of O(1).

*/

// NaiveSubArraySum ...
func NaiveSubArraySum(input []int, sum int) {
	for i := 0; i < len(input); i++ {
		//Here, we are trying to find sub-array starting from ith.
		s := 0
		for j := i; j < len(input); j++ {
			s += input[j]
			if s == sum {
				log.Printf("Sub-Array Found: Start: %d, End: %d\n", i, j)
				break
			}
		}
	}
}

/*
	2) HashMap Solution.

	This is again very easy and tricky approach.

	This will give us time complexity of O(n) and space complexity of O(n).

	Now, here the idea is, we need to store the sum of all possible sub-arrays.
	For Ex: Sum of first sub-array is 4.
	So, we will store this in map with Key as 4 and value as index of last element.
*/

// MapSubArraySum ...
func MapSubArraySum(input []int, sum int) {

	m := make(map[int]int, 0)

	current := 0
	for i := 0; i < len(input); i++ {

		current += input[i]
		if current == sum {
			log.Printf("Found Sub-Array: %d, %d\n", 0, i)
			continue
		}

		diff := current - sum
		lastIndex, ok := m[diff]
		if ok {
			log.Printf("Found Sub-Array: %d, %d\n", lastIndex+1, i)
		} else {
			m[current] = i
		}
	}
}
