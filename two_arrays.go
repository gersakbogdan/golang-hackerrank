package main

import (
    "fmt"
    "bufio"
    "os"
)

func sort(arr []int) []int {
    l := len(arr)

    for i := 1; i < l; i++ {
        x := arr[i]
        j := i

        for j > 0 && arr[j - 1] > x {
            arr[j] = arr[j - 1]
            j--
        }
        arr[j] = x
    }

    return arr
}

func main() {
    var t, n, k, temp int

    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &t)

    for t > 0 {
        fmt.Fscan(io, &n, &k)

        sum := 0

        a := make([]int, n)
        b := make([]int, n)

        for i := 0; i < n; i++ {
            fmt.Fscan(io, &temp)
            a[i] = temp
            sum += temp
        }

        for i := 0; i < n; i++ {
            fmt.Fscan(io, &temp)
            b[i] = temp
            sum += temp
        }

        a = sort(a)
        b = sort(b)

        if sum >= (n * k) {
            ynr1, ynr2 := 0, 0
            for i := 0; i < n; i++ {
                if (a[i] + b[i] >= k) {
                    ynr1++
                }
                if (a[n - i - 1] + b[i] >= k) {
                    ynr2++
                }
            }

            if ynr2 == n || ynr1 == n {
                fmt.Println("YES")
            } else {
                fmt.Println("NO")
            }
        } else {
            fmt.Println("NO")
        }

        t--
    }
}
