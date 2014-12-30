package main

import "fmt"

func main() {
    var n, p uint64

    fmt.Scan(&n)

    for n > 0 {
        fmt.Scan(&p)
        fmt.Println(p + 1)
        n--
    }
}
