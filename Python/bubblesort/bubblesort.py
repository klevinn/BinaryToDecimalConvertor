# Simple CLI Bubble Sort

def bubbleSort(arr):
    n = len(arr)
    for _ in range(n):
        for i in range(n-1):
            if arr[i] > arr[i+1]:
                arr[i], arr[i+1] = arr[i+1], arr[i]

def revBubbleSort(arr):
    n = len(arr)
    for _ in range(n):
        for i in range(n-1):
                if arr[i] < arr[i+1]:
                    arr[i], arr[i+1] = arr[i+1], arr[i]

def main():
    arr = [5,12,42,1,5,7,8]
    bubbleSort(arr)
    print(f"The sorted array is {arr}")
    revBubbleSort(arr)
    print(f"The reverse sorted array is {arr}")

main()