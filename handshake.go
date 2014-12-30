package main

import "fmt"

func main() {
	var nr, n int64

	fmt.Scan(&n)

	for n > 0 {
		fmt.Scan(&nr)

		fmt.Println(nr * (nr - 1) / 2)
		n--
	}
}