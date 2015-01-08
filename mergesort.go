package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

func mergeSort(arr []int) []int {
    l := len(arr)

    if l <= 1 {
        return arr
    }

    left := make([]int, 0)
    right := make([]int, 0)
    m := ceilDiv(l, 2)

    left = append(left, arr[0:m]...)
    right = append(right, arr[m:l]...)

    nLeft := mergeSort(left)
    nRight := mergeSort(right)

    return merge(nLeft, nRight)
}

func ceilDiv(a, b int) int {
    return int(math.Ceil(float64(a) / float64(b)))
}

func merge(left, right []int) []int {
    li := 0
    ri := 0
    ll := len(left)
    rl := len(right)

    mergeArr := make([]int, 0)

    for li < ll && ri < rl {
        if left[li] < right[ri] {
            mergeArr = append(mergeArr, left[li])
            li++
        } else {
            mergeArr = append(mergeArr, right[ri])
            ri++
        }
    }
    for li < ll {
        mergeArr = append(mergeArr, left[li])
        li++
    }
    for ri < rl {
        mergeArr = append(mergeArr, right[ri])
        ri++
    }

    return mergeArr
}

func main() {
    var n int

    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &n)

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(io, &arr[i])
    }

    fmt.Println(mergeSort(arr))
}
