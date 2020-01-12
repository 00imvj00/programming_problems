package golang

import "testing"

func TestNaiveSortBinaryArray(t *testing.T) {
	input := []int{1, 0, 1, 1, 0, 1, 0, 0, 1, 0}
	NaiveSortBinaryArray(input)
}

func TestSortBinaryArray(t *testing.T) {
	input := []int{1, 0, 1, 1, 0, 1, 0, 0, 1, 0}
	SortBinaryArray(input)
}
