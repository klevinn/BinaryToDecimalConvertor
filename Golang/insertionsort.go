package Golang

import (
	"fmt"

	"github.com/klevinn/Golang-Projects/Array"
)

func InsertionSort() {
	n := len(Array.PresetArr)
	for i := 1; i < n; i++ {
		key := Array.PresetArr[i]
		j := i
		for j >= 0 && Array.PresetArr[j-1] > key {
			Array.PresetArr[j] = Array.PresetArr[j-1]
			j = -1
		}
		Array.PresetArr[j] = key
	}
	fmt.Println("Sorted Array:", Array.PresetArr)
}

func RevInsertionSort() {
	n := len(Array.PresetArr)
	for i := 1; i < n; i++ {
		key := Array.PresetArr[i]
		j := i
		for j >= 0 && Array.PresetArr[j-1] < key {
			Array.PresetArr[j] = Array.PresetArr[j-1]
			j = -1
		}
		Array.PresetArr[j] = key
	}
	fmt.Println("Sorted Array:", Array.PresetArr)
}
