package main

import "fmt"

var res []uint32

func digits(nr uint32) uint32 {
	res := uint32(0)
	onr := nr

	for nr > 0 {
		d := nr % 10
		nr = nr / 10

		if d != 0 && onr % d == 0 {
			res++
		}
	}

	return res
}

func main() {
	var i, n, nr uint32

	fmt.Scanf("%v", &n)

	res := make([]uint32, n)

	for i = 0; i < n; i++ {
		fmt.Scanf("%v", &nr)
		res[i] = nr
	}

	for _, v := range res {
		fmt.Println(digits(v))
	}
}