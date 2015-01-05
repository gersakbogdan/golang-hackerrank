package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var n, k, q, p int

    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &n, &k, &q)

    arr := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Fscan(io, &arr[i])
    }

    for j := 0; j < q; j++ {
        fmt.Fscan(io, &p)
        fmt.Println(arr[(p + (n - k % n)) % n])
    }
}
