package Golang

func MergeSortedList(left []int, right []int) []int {
	leftInd := 0
	rightInd := 0
	result := []int{}
	for leftInd < len(left) && rightInd < len(right) {
		if left[leftInd] < right[rightInd] {
			result = append(result, left[leftInd])
			leftInd++
		} else {
			result = append(result, right[rightInd])
			rightInd++
		}
	}
	result = append(result, left[leftInd:]...)
	result = append(result, right[rightInd:]...)
	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		mid := len(arr) / 2
		lefthalf := MergeSort(arr[:mid])
		righthalf := MergeSort(arr[mid:])
		return MergeSortedList(lefthalf, righthalf)
	}
}
