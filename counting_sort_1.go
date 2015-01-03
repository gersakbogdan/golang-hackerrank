package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func print(arr []int) {
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}

func main() {
	buf := bufio.NewReader(os.Stdin)

	buf.ReadString('\n')
	line2, _ := buf.ReadString('\n')

	arr := make([]int, 100)

	for _, w := range strings.Fields(line2) {
		temp, _ := strconv.Atoi(w)
		arr[temp]++
	}

	print(arr)
}