package Golang

import (
	"fmt"
)

func PrintPresetArray() {
	arr := [...]int{5, 12, 42, 1, 5, 7, 8} // declare array
	fmt.Println("Array:", arr)
}

func BubbleSort() {
	arr := [...]int{5, 12, 42, 1, 5, 7, 8} // declare array
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				store := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = store
			}
		}
	}
	fmt.Println("Sorted Array:", arr)
}

func RevBubbleSort() {
	arr := [...]int{5, 12, 42, 1, 5, 7, 8} // declare array
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] < arr[j+1] {
				store := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = store
			}
		}
	}
	fmt.Println("Reverse Sorted Array:", arr)
}
