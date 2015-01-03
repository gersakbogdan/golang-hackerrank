package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	var a [100]int
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(io, &v)
		a[v]++
	}

    print(a[:])
}

func print(slice []int) {
	for i, l := 0, len(slice)-1; i < l; i++ {
		fmt.Print(slice[i], " ")
	}

	fmt.Println(slice[len(slice)-1])
}
