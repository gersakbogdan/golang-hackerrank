package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

func partition(lo, hi int, arr []int) []int {
    poz := lo

    for i := lo; i < hi; i++ {
        if arr[i] < arr[hi] {
            if i != poz {
                arr[i], arr[poz] = arr[poz], arr[i]
            }
            poz++

        }
    }
    arr[hi], arr[poz] = arr[poz], arr[hi]

    if poz - 1 > lo {
        partition(lo, poz - 1, arr)
    }
    if poz + 1 < hi {
        partition(poz + 1, hi, arr)
    }

    return arr
}

func quicksort(arr []int) []int {
    l := len(arr)
    hi := l - 1

    return partition(0, hi, arr)
}

func abs(n int) int {
    return int(math.Abs(float64(n)))
}

func main() {
    var n int;

    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &n)

    arr := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Fscan(io, &arr[i])
    }

    quicksort(arr)

    diff := arr[len(arr) - 1] - arr[0]

    for i := 0; i < len(arr) - 1; i++ {
        if abs(arr[i] - arr[i+1]) < diff {
            diff = abs(arr[i] - arr[i+1])
        }
    }

    for i := 0; i < len(arr) - 1; i++ {
        if abs(arr[i] - arr[i+1]) == diff {
            fmt.Print(arr[i], arr[i + 1], " ")
        }
    }

    fmt.Println()
}
