package main

import "fmt"

func insertionSortTime(arr []uint64) int {
    l := len(arr)
    op := 0

    for i := 1; i < l; i++ {
        x := arr[i]
        j := i

        for j > 0 && arr[j - 1] > x {
            arr[j] = arr[j-1]
            j--

            op++
        }

        arr[j] = x
    }

    return op
}

func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]uint64, n)

    for i:=0; i<n; i++ {
        fmt.Scan(&arr[i])
    }

    fmt.Println(insertionSortTime(arr))
}
