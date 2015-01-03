package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var ar [100]int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		ar[num]++
	}

	for i := 0; i < 100; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()

}