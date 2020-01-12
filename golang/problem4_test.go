package golang

import "testing"

func TestFindOccurences(t *testing.T) {
	input := []int{1, 2, 4, 7, 2, 3, 6, 3, 4, 5, 7, 3, 3, 4, 5}
	FindOccurences(input)
}
