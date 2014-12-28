package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	ch := make([]int, 26)

	for _, v := range s {
		i := int(v) - int('a')
		ch[i] = (ch[i] + 1) % 2
	}

	sum := 0
	for _, v := range ch {
		sum += v
	}

	if len(s) % 2 == sum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}