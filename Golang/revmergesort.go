package Golang

func RevMergeSortedList(left []int, right []int) []int {
	leftInd := 0
	rightInd := 0
	result := []int{}
	for leftInd < len(left) && rightInd < len(right) {
		if left[leftInd] > right[rightInd] {
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

func RevMergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		mid := len(arr) / 2
		lefthalf := RevMergeSort(arr[:mid])
		righthalf := RevMergeSort(arr[mid:])
		return RevMergeSortedList(lefthalf, righthalf)
	}
}
