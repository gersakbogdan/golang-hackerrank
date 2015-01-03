package main

import (
	"fmt"
	"bufio"
	"os"
)

func print(arr [100]int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < arr[i]; j++ {
			fmt.Print(i, " ")
		}
	}

	fmt.Println()
}

func main() {
	var v, n int
	var arr [100]int

	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	for n > 0 {
		fmt.Fscan(io, &v)
		arr[v]++
		n--
	}

	print(arr)
}