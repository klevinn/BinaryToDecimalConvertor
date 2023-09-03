package main

import (
	"fmt"

	"github.com/klevinn/Golang-Projects/Array"
	"github.com/klevinn/Golang-Projects/Golang"
)

func menu() {
	fmt.Println("1. Decimal To Binary / Binary To Decimal")
	fmt.Println("2. Bubble Sort / Reverse Bubble Sort")
	fmt.Println("3. Selection Sort / Reverse Selection Sort")
	fmt.Println("4. Exit")
}

func main() {
	menu()
	for {
		var choice int
		fmt.Print("Which function would you like to see? ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Please enter a valid option")
		} else {
			if choice == 1 {
				fmt.Println("1. Decimal To Binary")
				Golang.DecimalToBinary()
				fmt.Println("2. Binary To Decimal")
				Golang.BinaryToDecimal()
			} else if choice == 2 {
				Array.PrintPresetArray()
				fmt.Println("1. Bubble Sort")
				Golang.BubbleSort()
				fmt.Println("2. Reverse Bubble Sort")
				Golang.RevBubbleSort()
			} else if choice == 3 {
				Array.PrintPresetArray()
				fmt.Println("1. Selection Sort")
				Golang.SelectionSort()
				fmt.Println("2. Reverse Selection Sort")
				Golang.RevSelectionSort()
			} else if choice == 4 {
				break
			}
		}
	}
}
