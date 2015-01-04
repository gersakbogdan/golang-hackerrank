package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func print(arr []int) {
	acc := 0

	for i := 0; i < 100; i++ {
		acc += arr[i]
		fmt.Print(acc, " ")
	}

	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, 100)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err == nil {
			arr[num]++	
		}
	}

	print(arr)
}