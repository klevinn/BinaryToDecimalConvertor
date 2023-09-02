package golang

import (
	"fmt"
)

func bubbleSort() {
	arr := [10]int{5, 12, 42, 1, 5, 7, 8} // declare array
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[i] > arr[i+1] {
				store := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = store
			}
		}
	}
	fmt.Println("Sorted Array:", arr)
}
