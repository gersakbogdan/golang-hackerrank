package main

import "fmt"

func solveMeFirst(a, b uint32) uint32 {
	return a + b
}

func main() {
	var a, b, res uint32;

	fmt.Scanf("%v %v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}