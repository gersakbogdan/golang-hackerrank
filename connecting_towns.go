package main

import (
    "fmt"
    "math/big"
)

func main() {
    var n, t, r int64
    modulo := big.NewInt(1234567)

    fmt.Scan(&n)

    for n > 0 {
        nrr := big.NewInt(1)
        fmt.Scan(&t)

        for t > 1 {
            fmt.Scan(&r)
            nrr.Mul(big.NewInt(r), nrr)
            t--
        }

        fmt.Println(nrr.Mod(nrr, modulo) )
        n--
    }
}
