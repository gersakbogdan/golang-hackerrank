package main

import "fmt"

func dash6(arr []int) []int {
	poz := 0
	l := len(arr)

	for i := 1; i < l; i++ {
		j := i;

		for j < l && arr[j] > arr[poz] {
			j++
		}

		if j < l {
			x := arr[j]

			for j > poz {
				arr[j] = arr[j-1]
				j--
			}

			poz++
			arr[j] = x
		}
	}

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

	print(dash6(arr))
}