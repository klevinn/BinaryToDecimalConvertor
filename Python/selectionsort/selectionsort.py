# Simple CLI Selection Sort

def selectionSort(arr):
    n = len(arr)
    for i in range(n):
        lowest = i
        for j in range(i+1,n):
            if arr[j] < arr[lowest]:
                lowest = j

        arr[i], arr[lowest] = arr[lowest], arr[i]

def revSelectionSort(arr):
    n = len(arr)
    for i in range(n):
        lowest = i
        for j in range(i+1,n):
            if arr[j] > arr[lowest]:
                lowest = j

        arr[i], arr[lowest] = arr[lowest], arr[i]

def main():
    arr = [5,12,42,1,5,7,8]
    selectionSort(arr)
    print(f"The sorted array is {arr}")
    revSelectionSort(arr)
    print(f"The reverse sorted array is {arr}")

main()