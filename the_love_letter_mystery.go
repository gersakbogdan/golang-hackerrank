package main

import (
    "fmt"
    "math"
)

func operations(s string) int {
    l := len(s)
    hl := l / 2
    op := 0

    for i := 0; i < hl; i++ {
        if s[i] != s[l-i-1] {
            op += abs(int(s[i]) - int(s[l-i-1]))
        }
    }

    return op
}

func abs(n int) int {
    return int(math.Abs(float64(n)))
}

func main() {
    var n int
    var s string

    fmt.Scan(&n)

    for n > 0 {
        fmt.Scan(&s)
        fmt.Println(operations(s))
        n--
    }
}
