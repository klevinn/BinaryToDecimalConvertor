package Golang

import (
	"fmt"

	"github.com/klevinn/Golang-Projects/Array"
)

func BubbleSort() {
	n := len(Array.PresetArr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if Array.PresetArr[j] > Array.PresetArr[j+1] {
				store := Array.PresetArr[j]
				Array.PresetArr[j] = Array.PresetArr[j+1]
				Array.PresetArr[j+1] = store
			}
		}
	}
	fmt.Println("Sorted Array:", Array.PresetArr)
}

func RevBubbleSort() {
	n := len(Array.PresetArr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if Array.PresetArr[j] < Array.PresetArr[j+1] {
				store := Array.PresetArr[j]
				Array.PresetArr[j] = Array.PresetArr[j+1]
				Array.PresetArr[j+1] = store
			}
		}
	}
	fmt.Println("Reverse Sorted Array:", Array.PresetArr)
}
