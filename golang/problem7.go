package golang

import "log"

// SortBasedP7Solution ...
func SortBasedP7Solution(input []int) {

	pivotElement := 1
	startPivotIndex := 0
	endPivotIndex := len(input) - 1
	start := 0

	for start <= endPivotIndex {

		if input[start] < pivotElement {
			Swap(input, start, startPivotIndex)
			startPivotIndex++
			start++
		} else if input[start] > pivotElement {
			Swap(input, start, endPivotIndex)
			endPivotIndex--
		} else {
			start++
		}

	}

	log.Printf("OutPut: %#v\n", input)
}
