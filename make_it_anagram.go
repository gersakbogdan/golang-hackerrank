package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

func abs(v int) int {
    return int(math.Abs(float64(v)))
}

func main() {
    var s1, s2 string

    deletions := 0
    alphabet := make([]int, 26)

    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &s1, &s2)

    for _, char := range s1 {
        alphabet[int(char) - int('a')]++
    }
    for _, char := range s2 {
        alphabet[int(char) - int('a')]--
    }

    for i := 0; i < 26; i++ {
        deletions += abs(alphabet[i])
    }

    fmt.Println(deletions)
}
