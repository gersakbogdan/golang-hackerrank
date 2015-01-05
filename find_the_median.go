package main

import (
    "fmt"
    "os"
    "bufio"
)

func partition(lo, hi int, arr []int) int {
    piv := hi
    poz := lo

    for i := lo; i < hi; i++ {
        if arr[i] < arr[piv] {
            if i != poz {
                arr[i], arr[poz] = arr[poz], arr[i]
            }
            poz++
        }
    }
    arr[poz], arr[piv] = arr[piv], arr[poz]

    return poz
}

func median(lo, hi int, arr []int) int {
    n := len(arr)
    poz := partition(lo, hi, arr)

    if poz == n / 2 {
        return arr[poz]
    } else if poz < n / 2 {
        return median(poz + 1, hi, arr)
    } else {
        return median(lo, poz - 1, arr)
    }
}

func main() {
    var n int
    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &n)

    arr := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Fscan(io, &arr[i])
    }

    fmt.Println(median(0, n - 1, arr))
}
