package golang

import (
	"log"
)

/*

	1) Naive Approach

	Here, we need to find a pair, so what basically we are doing is for every element in array,
	we are trying to find another element and then we are creating pair, then we are summing the pair
	and see if this is the pair we are looking for?

	In the naive solution, we are pairing each element of array with each other element of array.

	For that we have to go through array n^2 times.

	So, because of that reason, the time complexity of this function will be O(n^2).

	Now, we are not allocating any extra space for this, so, space complexity is going to be O(1).

*/

//NaiveFindPair ...
func NaiveFindPair(input []int, sum int) {

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			tempSum := input[i] + input[j]
			if tempSum == sum {
				log.Printf("Pair: (%d, %d)\n", input[i], input[j])
			}
		}
	}

}

/*

	2) Sorting Approach

	Well, we know that array is not sorted. And we know that we want to find sum of pair which is 10 in this case.
	And we go through the array, we have this equation in our hand, i + x = 10; where i is the current element.

	So, we need to find x in the equation.

	Now, we know that sum should not be more than 10. And we have sorted array in hand in ascending order.

	So, Now, if we start from the beginning of the array, we have smallest element in our hand, we set this as i in equation.

	Now, we will take last element which is the largest element in the array as x.

	Then we will try to sum them and see if we are getting the expected sum or not.

	If calculated sum is greater than the sum that we are looking for, in that case we will know that, we need to reduce the value of one of the
	element.

	Now i is already smallest, so, we will reduce the value of x, by selecting previous largest element from the list.

	Now, if the calculated sum is less than given sum, then we will know we have to increase the value of one of the element.

	So, for that we will increase the value of i, by selecting the next large element from i.

	We will keep doing this activity, intil we don't cross i and x.

	Please refer code to get more explanation.

*/

// SortFindPair ...
func SortFindPair(input []int, sum int) {

	//First step we need to do, is sort the array, which will incure O(nlogn)
	// In this case we will use quick sort for sorting array.

	QuickSort(input)

	//now, we will maintain i and x variables.
	i := 0              //Smallest element
	x := len(input) - 1 //largest element

	//Now, we will go through the array, until i >= x

	for x > i {

		c := input[i] + input[x]

		if c > sum {
			x--
			continue
		}

		if c < sum {
			i++
			continue
		}

		log.Printf("Found Pair: (%d, %d)\n", input[i], input[x])
		x--
		i++
		continue
	}

}

/*
	3) HashMap Approach

	There is one more way, we can solve this using hashmap.

	This will give us O(n) time complexity,
	As we are storing the elements in hashmap, we will incur O(n) space complexity.

	Again, let's convert problem into math: i + x = sum;
	Now, we have i, sum, we need to find x.

	So, can we also write this equation like this? x = sum - i;

	Yes, we can.

	So, now the solution is easy, because, we just need to search for x in given array.
	what we can do is, we can store all the elements in hashmap.

	Then for each i, we will calculate x. Now, if calculated x is in the hashmap, then that is out pair.

	Simple. !!!
*/

// MapFindPair ...
func MapFindPair(input []int, sum int) {
	m := make(map[int]bool, 0)

	for i := 0; i < len(input); i++ {

		x := sum - input[i]
		_, ok := m[x]

		if ok {
			log.Printf("Found Pair: (%d, %d)\n", input[i], x)
		} else {
			m[input[i]] = true
		}

	}

}
