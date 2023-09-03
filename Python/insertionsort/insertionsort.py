# Simple CLI Insertion Sort

def insertionSort(arr):
    n = len(arr)
    for i in range(1,n):
        value = arr[i]
        current = i
        while (current > 0) and (arr[current-1] > value):
            arr[current] = arr[current-1]
            current -= 1

        
        arr[current] = value

def revInsertionSort(arr):
    n = len(arr)
    for i in range(1,n):
        value = arr[i]
        current = i
        while (current > 0) and (arr[current-1] < value):
            arr[current] = arr[current-1]
            current -= 1
        
        arr[current] = value

def main():
    arr = [5,12,42,1,5,7,8]
    insertionSort(arr)
    print(f"The sorted array is {arr}")
    revInsertionSort(arr)
    print(f"The reverse sorted array is {arr}")

main()