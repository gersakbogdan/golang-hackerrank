package main

import (
    "fmt"
    "math/big"
)

func fiboModified(a, b, nth int64) *big.Int {
    var i int64

    tn  := big.NewInt(a)
    tn1 := big.NewInt(b)

    for i = 2; i < nth; i = i + 2 {
        temp := new(big.Int)

        temp.Mul(tn1, tn1)
        tn.Add(temp, tn)

        temp.Mul(tn, tn)
        tn1.Add(temp, tn1)
    }

    if nth % 2 == 1 {
        return tn
    }
    return tn1
}

func main() {
    var a, b, nth int64
    fmt.Scan(&a, &b, &nth)
    fmt.Println(fiboModified(a, b, nth))
}
