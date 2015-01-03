package main

import "fmt"

func partition(lo int, piv int, arr []int) int {
	op := 0
	is := lo

	for i := lo; i < piv; i++ {
		if arr[i] < arr[piv] {
			if i != is {
				arr[i], arr[is] = arr[is], arr[i]
			}
			op++
			is++
		}
	}

	arr[is], arr[piv] = arr[piv], arr[is]
	op++

	if is - 1 > lo {
		op += partition(lo, is - 1, arr)
	}
	if is + 1 < piv {
		op += partition(is + 1, piv, arr)
	}

	return op
}

func quickSortTime(arr []int) int {
	l := len(arr)
	piv := l - 1

	return partition(0, piv, arr)
}

func insertionSortTime(arr []int) int {
    l := len(arr)
    op := 0

    for i := 1; i < l; i++ {
        x := arr[i]
        j := i

        for j > 0 && arr[j - 1] > x {
            arr[j] = arr[j-1]
            j--

            op++
        }

        arr[j] = x
    }

    return op
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	arr2 := make([]int, 0)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	arr2 = append(arr2, arr...)

	iop := insertionSortTime(arr2)
	qop := quickSortTime(arr)

	fmt.Println(iop - qop)
}