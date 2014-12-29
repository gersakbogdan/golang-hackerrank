package main

import "fmt"

var n int

func printArr(arr []uint64, n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("%v ", arr[i])
    }
    fmt.Println()
}

func insertionSort(arr []uint64, n int) {
    i := n - 1
    j := i
    x := arr[i]

    for j > 0 && arr[j - 1] > x {
        arr[j] = arr[j -1]
        j--
        printArr(arr, n)
    }
    arr[j] = x
    printArr(arr, n)
}

func main() {
    fmt.Scan(&n)
    arr := make([]uint64, n)

    for i:=0; i<n; i++ {
        fmt.Scan(&arr[i])
    }

    insertionSort(arr, n)
}
