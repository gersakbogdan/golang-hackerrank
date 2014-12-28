package main

import (
	"fmt"
	"math"
)

func main() {
	var t, n int
	res := make([]int, 0)

	fmt.Scanf("%v", &t)

	for t > 0 {
		fmt.Scanf("%v", &n)
		res = append(res, int(math.Pow(2, float64(int((n+1)/2 + 1)))) - 1 - (n % 2) )
		t--
	}

	for _, v := range res {
		fmt.Println(v)
	}
}