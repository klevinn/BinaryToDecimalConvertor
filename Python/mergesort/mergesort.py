# Simple CLI Merge Sort
def mergeSortedList(left, right):
    leftInd = 0
    rightInd = 0
    result = []
    
    while (leftInd < len(left) and (rightInd < len(right))):
        if (left[leftInd] < right[rightInd]):
            result.append(left[leftInd])
            leftInd += 1
        else:
            result.append(right[rightInd])
            rightInd += 1

    result += left[leftInd:]
    result += right[rightInd:]
    return result

def mergeSort(arr):
    if (len(arr) <= 1):
        return arr
    else:
        mid = len(arr) // 2

        lefthalf = mergeSort(arr[:mid])
        righthalf = mergeSort(arr[mid:])

        newList = mergeSortedList(lefthalf, righthalf)

        return newList
    
def main():
    arr = [5,12,42,1,5,7,8]
    sorterdarr = mergeSort(arr)
    print(f"The sorted array is {sorterdarr}")

main()