package main

import "fmt"

func fibModified(a, b, nth uint64) uint64 {

    for nth > 0 {
        a = b*b + a
        b = a*a + b
        nth--
        nth--
    }

    return b
}

func main() {
    var a, b, nth uint64
    fmt.Scan(&a, &b, &nth)

    fmt.Println(fibModified(a, b, nth))
}
