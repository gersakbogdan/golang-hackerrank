package main

import "fmt"

func partition(lo int, piv int, arr []int) int {
	is := lo

	for i := lo; i < piv; i++ {
		if arr[i] < arr[piv] {
			if i != is {
				arr[i], arr[is] = arr[is], arr[i]
			}

			is++
		}
	}

	arr[is], arr[piv] = arr[piv], arr[is]
	print(arr)

	if is - 1 > lo {
		partition(lo, is - 1, arr)
	}
	if is + 1 < piv {
		partition(is + 1, piv, arr)
	}

	return is
}

func quickSort(arr []int) []int {
	l := len(arr)
	piv := l - 1

	partition(0, piv, arr)

	return arr
}

func print(arr []int) {
	l := len(arr)

	for i := 0; i < l; i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	quickSort(arr)
}