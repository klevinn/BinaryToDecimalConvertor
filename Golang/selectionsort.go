package Golang

import (
	"fmt"

	"github.com/klevinn/Golang-Projects/Array"
)

func SelectionSort() {
	n := len(Array.PresetArr)
	for i := 0; i < n; i++ {
		lowest := i
		for j := i + 1; j < n; j++ {
			if Array.PresetArr[j] < Array.PresetArr[lowest] {
				lowest = j
			}
		}

		store := Array.PresetArr[i]
		Array.PresetArr[i] = Array.PresetArr[lowest]
		Array.PresetArr[lowest] = store
	}

	fmt.Println("Sorted Array:", Array.PresetArr)
}

func RevSelectionSort() {
	n := len(Array.PresetArr)
	for i := 0; i < n; i++ {
		lowest := i
		for j := i + 1; j < n; j++ {
			if Array.PresetArr[j] > Array.PresetArr[lowest] {
				lowest = j
			}
		}

		store := Array.PresetArr[i]
		Array.PresetArr[i] = Array.PresetArr[lowest]
		Array.PresetArr[lowest] = store
	}

	fmt.Println("Reverse Sorted Array:", Array.PresetArr)
}
