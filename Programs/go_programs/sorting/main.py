import random
import numpy as np

arr = []

def createRandomArray(n):
    global arr
    arr = np.random.randint(0, n, size=n)

def mergeSort(arr):
    # check if len of arr > 1 and find mid, left and right arrays
    if len(arr) > 1:
        mid = len(arr) // 2
        left = arr[:mid]
        right = arr[mid:]

        # recurse on left and right 
        mergeSort(left)
        mergeSort(right)

        # i for left and j for right 
        i = 0
        j = 0
        # k for the main array arr
        k = 0
        
        # replace arr values as per sorted values
        while i < len(left) and j < len(right):
            # if val on left <= val on right, replace arr[k] with left val
            if left[i] <= right[j]:
              arr[k] = left[i]
              i += 1
            else: # replace arr[k] with val on right 
                arr[k] = right[j]
                j += 1
            k += 1
        # if there are any remaining values (for example, in case right and left arrays have values left)
        while i < len(left):
            arr[k] = left[i]
            i += 1
            k += 1
        while j < len(right):
            arr[k]=right[j]
            j += 1
            k += 1

if __name__ == "__main__":
    n = 10000000
    createRandomArray(n)
    mergeSort(arr)
    #print(arr)