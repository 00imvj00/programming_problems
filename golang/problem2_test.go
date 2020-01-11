package golang

import "testing"

func TestNaiveSubArraySumZero(t *testing.T) {
	input := []int{4, 2, -3, -1, 0, 4}
	sum := 0
	NaiveSubArraySum(input, sum)
}

func TestMapSubArraySum(t *testing.T) {
	input := []int{4, 2, -3, -1, 0, 4}
	sum := 0
	MapSubArraySum(input, sum)
}
