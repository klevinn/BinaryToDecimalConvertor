# Merge Sort

This is a simple command line interface (CLI) program that can perform Merge Sort on an array of numbers. There is a both a Python and Golang version in this repository

## Code Explanation

The `mergeSort()` function sorts an array of numbers in ascending order using bubble sort while the `revMergeSort()` function sorts it in a descending order. Each Python file has its own version of Merging Algorithm to accommodate for the Reverse Sorting

For Golang since camelCase functions are package-private functions the functions were renamed to `MergeSort` and `RevMergeSort`

Merge Sort Function
    Best time complexity: O(n log n)
    Worst time complexity: O(n log n)
    Average time complexity: O(n log n)

Merge Sort breaks the list into 2 halves, then recursively calls itself on the 2 halves until the list is 1 element
Then we merge the 2 halves together all the way till the list is sorted

Demonstration

Initial = [5,1,2,4]
First Split = [5,1] & [2,4]
Second Split = [5], [1], [2], [4]

Merge together = [1,5] , [2,4]
Merge Together = [1,2,4,5]

Keeps going through all elements of the list