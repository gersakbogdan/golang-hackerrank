package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var t, n, v, answer int
    io := bufio.NewReader(os.Stdin)

    fmt.Fscan(io, &t)

    for t > 0 {
        answer = 0
        fmt.Fscan(io, &n)

        for i := 0; i < n; i++ {
            fmt.Fscan(io, &v)

            times := uint64(i + 1) * uint64(n - i)
            if times % 2 == 1 {
                answer ^= v
            }
        }

        fmt.Println(answer)
        t--
    }
}
