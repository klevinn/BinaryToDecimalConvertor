# Bubble Sort

This is a simple command line interface (CLI) program that can perform Selection Sort on an array of numbers. There is a both a Python and Golang version in this repository

## Code Explanation

The `selectionSort()` function sorts an array of numbers in ascending order using bubble sort while the `revSelectionSort()` function sorts it in a descending order

For Golang since camelCase functions are package-private functions the functions were renamed to `SelectionSort()` and `RevSelectionSort()`

Selection Sort Function
    Best time complexity: O(n^2)
    Worst time complexity: O(n^2)
    Average time complexity: O(n^2)

Compare element with all elements till lowest is found and swaps with the first index

First to ensure we check through all the elements, we use the for i in range(n) loop:
Then we check through the rest of the elements using the for j in range(i+1,n) loop

Demonstration

Initial = [5,1,2]
First Loop : 
    lowest variable is assigned to index 1 or the value 5
    Since 1 is lesserr than 5, the index 2 is assigned to the lowest variable
Second Loop : 
    lowest variable stays assigned to index 2 as 1 is lesser than 2
After all comparisons with rest of the elements:
The element at index 1 (lowest variable) swaps with the first index:
[1,5,2]

This goes on for all elements