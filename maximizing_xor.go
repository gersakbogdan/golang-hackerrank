package main

import "fmt"

func main() {
	var l, r, max int

	fmt.Scanf("%v", &l)
	fmt.Scanf("%v", &r)

	for i := l; i <= r; i++ {
		for j := i; j <=r; j++ {
			xor := i ^ j

			if xor > max {
				max = xor
			}
		}
	}

	fmt.Println(max)
}