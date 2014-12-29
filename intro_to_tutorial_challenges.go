package main

import "fmt"

func binarySearch(arr []uint64, target uint64) int {
    lo := 0
    hi := len(arr) - 1

    for lo <= hi {
        mid := lo + (hi - lo) / 2

        if arr[mid] > target {
            hi = mid - 1
        } else if arr[mid] < target {
            lo = mid + 1
        } else {
            return mid
        }
    }

    return -1
}

func main() {
    var n int
    var v uint64

    fmt.Scan(&v, &n)

    arr := make([]uint64, n)

    for i:=0; i<n; i++ {
        fmt.Scan(&arr[i])
    }

    fmt.Println(binarySearch(arr, v))
}
