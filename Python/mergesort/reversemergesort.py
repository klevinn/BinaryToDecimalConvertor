# Simple CLI Reverse Merge Sort

def mergeSortedList(left, right):
    leftInd = 0
    rightInd = 0
    result = []

    while (leftInd < len(left) and (rightInd < len(right))):
        if (left[leftInd] > right[rightInd]):
            result.append(left[leftInd])
            leftInd += 1
        else:
            result.append(right[rightInd])
            rightInd += 1   

    result += left[leftInd:]
    result += right[rightInd:]
    return result

def revMergeSort(arr):
    if (len(arr) <= 1):
        return arr
    else:
        mid = len(arr) // 2

        lefthalf = revMergeSort(arr[:mid])
        righthalf = revMergeSort(arr[mid:])

        newList = mergeSortedList(lefthalf, righthalf)

        return newList
    
def main():
    arr = [5,12,42,1,5,7,8]
    sortedarr = revMergeSort(arr)
    print(f"The sorted array is {sortedarr}")

main()