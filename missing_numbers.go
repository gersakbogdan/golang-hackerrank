package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var n, t int
	
	a := make([]int, 10001)
	b := make([]int, 10001)

	io := bufio.NewReader(os.Stdin)

	fmt.Fscan(io, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(io, &t)
		a[t]++
	}

	fmt.Fscan(io, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(io, &t)
		b[t]++
	}

	// Missing numbers
	for i := 0; i < 10001;  i++ {
		if b[i] - a[i] > 0 {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()
}