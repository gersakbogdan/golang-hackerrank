package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func print(arr [][]string) {
	fmt.Println(strings.Replace(strings.Replace(fmt.Sprint(arr), "[", "", -1), "]", "", -1))
}

func main() {
	var n int
	fmt.Scan(&n)

	m := n / 2

	arr := make([][]string, 100)
	scanner := bufio.NewScanner(os.Stdin)

	index := 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		i, _ := strconv.Atoi(fields[0])

		if index < m {
			arr[i] = append(arr[i], "-")
		} else {
			arr[i] = append(arr[i], fields[1])
		}

		index++
	}
	
	print(arr)
}