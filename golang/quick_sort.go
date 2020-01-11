package golang

// QuickSort ...
func QuickSort(data []int) {
	qs(data, 0, len(data)-1)
}

func qs(data []int, lb, ub int) {

	if lb < ub {
		pivotPos := partition(data, lb, ub)
		qs(data, lb, pivotPos-1)
		qs(data, pivotPos+1, ub)
	}

}

func partition(data []int, lb, ub int) int {

	pivot := data[ub]
	pivotIndex := lb

	for i := lb; i < ub; i++ {
		if data[i] <= pivot {
			Swap(data, i, pivotIndex)
			pivotIndex++
		}
	}

	Swap(data, pivotIndex, ub)
	return pivotIndex
}

// Swap ...
func Swap(data []int, i, j int) {
	temp := data[i]
	data[i] = data[j]
	data[j] = temp
}
