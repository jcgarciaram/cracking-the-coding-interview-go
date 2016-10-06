package main

import (
    "fmt"
    "math"
)

func main() {
    arr1 := []int{1, 3, 15, 11, 2}
    arr2 := []int{23, 127, 235, 19, 8}
    //fmt.Println(arr1, arr2)
    fmt.Println(smallestDiff(arr1, arr2))
}

func quickSort(arr []int, left, right int) []int {
    var index int
    index, arr = partition(arr, left, right)
    //fmt.Println(index)
    if left < index-1 {
        arr = quickSort(arr, left, index-1)
    }
    if index < right {
        arr = quickSort(arr, index, right)
    }
    return arr
}
func partition(arr []int, left, right int) (int, []int) {
    pivot := arr[(left+right)/2]
    //fmt.Println(pivot)
    for left <= right {
        //fmt.Println("Left: ", left, "Right: ", right) 
        for arr[left] < pivot {
            left++
        }
        for arr[right] > pivot {
            right--
        }
        if left <= right {
            tempLeft := arr[left]
            arr[left] = arr[right]
            arr[right] = tempLeft
            left++
            right--
        }
    }
    return left, arr
}

func smallestDiff(arr1, arr2 []int) int {
    
    arr1 = quickSort(arr1, 0, len(arr1)-1)

    minDiff := math.MaxInt32
    for _, val := range arr2 {
        lDiff, rDiff := binaryDiff(val, arr1)
        if lDiff == 0 || rDiff == 0 {
            return 0
        }
        if lDiff < minDiff {
            minDiff = lDiff
        }
        if rDiff < minDiff {
            minDiff = rDiff
        }
    }
    
    return minDiff
}

func binaryDiff(val int, arr []int) (int, int) {
    
    if len(arr) == 2 {
        if val < arr[0] {
            return math.MaxInt32, arr[0]-val
        } else if val > arr[1] {
            return val-arr[1], math.MaxInt32
        } else {
            return val-arr[0], arr[1]-val
        }
    }
    if len(arr) == 1 {
        if val < arr[0] {
            return math.MaxInt32, arr[0]-val
        } else if val > arr[0] {
            return val-arr[0], math.MaxInt32
        } else {
            return 0, 0
        }
    }
    midIndex := len(arr)/2
    if val > arr[midIndex] {
        return binaryDiff(val, arr[midIndex+1:])
    } else if val < arr[midIndex] {
        return binaryDiff(val, arr[:midIndex])
    }
    return 0,0
}
    
