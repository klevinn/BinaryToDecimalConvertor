# Bubble Sort

This is a simple command line interface (CLI) program that can perform Bubble Sort on an array of numbers. There is a both a Python and Golang version in this repository

## Code Explanation

The `bubbleSort()` function sorts an array of numbers in ascending order using bubble sort while the `revBubbleSort` function sorts it in a descending order

For Golang since camelCase functions are package-private functions the functions were renamed to `BubbleSort` and `RevBubbleSort`

Bubble Sort Function
    Best time complexity: O(n)
    Worst time complexity: O(n^2)
    Average time complexity: O(n^2)

Greater Element Bubbles up to the top

First we check through all the elements using the for _ in range(n) loop
Then we check through all the elements using the for i in range(n-1) loop

Demonstration

Initial : [5,1,2]
First Loop : [1,5,2]
Second Loop : [1,2,5]