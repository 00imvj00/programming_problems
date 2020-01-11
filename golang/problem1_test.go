package golang

import "testing"

func TestNaiveFindPair(t *testing.T) {
	input := []int{8, 7, 2, 5, 3, 1}
	sum := 10
	NaiveFindPair(input, sum)
}

func TestSortFindPair(t *testing.T) {
	input := []int{8, 7, 2, 5, 3, 1}
	sum := 10
	SortFindPair(input, sum)
}

func TestMapFindPair(t *testing.T) {
	input := []int{8, 7, 2, 5, 3, 1}
	sum := 10
	MapFindPair(input, sum)
}
