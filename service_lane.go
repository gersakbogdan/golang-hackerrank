package main

import (
    "fmt"
    "bufio"
    "os"
)

func largeVehicle(start, end int, freeway []int) int {
    min := 3 // truck

    for i := start; i <= end; i++ {
        if freeway[i] < min {
            min = freeway[i]
        }
    }

    return min
}

func main() {
    var n, t, start, end int

    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &n, &t)

    freeway := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(io, &freeway[i])
    }

    for t > 0 {
        fmt.Fscan(io, &start, &end)
        fmt.Println(largeVehicle(start, end, freeway))
        t--
    }
}
