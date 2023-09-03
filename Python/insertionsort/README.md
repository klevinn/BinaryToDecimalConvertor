# Insertion Sort

This is a simple command line interface (CLI) program that can perform Insertion Sort on an array of numbers. There is a both a Python and Golang version in this repository

## Code Explanation

The `insertionSort()` function sorts an array of numbers in ascending order using bubble sort while the `revinsertionSort()` function sorts it in a descending order

For Golang since camelCase functions are package-private functions the functions were renamed to `InsertionSort` and `RevInsertionSort`

Insertion Sort Function
    Best time complexity: O(n)
    Worst time complexity: O(n^2)
    Average time complexity: O(n^2)

Similar to Selection, in terms of comparing through all elements in the list. However, Insertion works backwards and instead of storing the index, swaps the element with the previous element if the previous element is greater

The for i in range (1,n) starts from 1 due to us minusing the index by 1 if swap, hence if start from 0 will return index error / redundent code as the while loop will never be entered
The while loop basically checks the current element with the previous element, if the previous element is greater than the current element, we replace the current element with the previous element. The -= 1 is to continuously move down the list
arr[current] = value places the replaced value in the current index

Demonstration

Initial = [5,1,2]
First loop : 
    current = 1
    1 is greater than 5, so the current element gets replaced
    [5,5,2]
    The current value minuses one so = 0, exitting the loop
    Replacing the current with the replaced element
    [1,5,2]

Keeps going through all elements of the list