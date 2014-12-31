package main

import "fmt"

func insertionSort(arr []uint64) {
    l := len(arr)

    for i := 1; i < l; i++ {
        x := arr[i]
        j := i

        for j > 0 && arr[j - 1] > x {
            arr[j] = arr[j-1]
            j--
        }
        arr[j] = x
    }

    print(arr)
}

func print(arr []uint64) {
    l := len(arr)
    for i := 0; i < l; i++ {
        fmt.Printf("%v ", arr[i])
    }
    fmt.Println()
}

func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]uint64, n)

    for i:=0; i<n; i++ {
        fmt.Scan(&arr[i])
    }

    insertionSort(arr)
}
