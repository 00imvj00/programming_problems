package golang

import "log"

/*
	1) HasMap Approach

	The idea is simple.

	As we are iterating through all the elements in array,
	we will first see if is there any element in the map.

	If there is in that case we will just update the occurence count for that element.

	If there is not, in that case, we will just add that element in that map.

	At the end we will just print the map.

	Simple !!!
*/

// FindOccurences ...
func FindOccurences(input []int) {
	m := make(map[int]int, 0)

	for _, value := range input {
		o, ok := m[value]
		if ok {
			m[value] = o + 1
		} else {
			m[value] = 1
		}
	}

	log.Printf("OutPut: %#v\n", m)
}
