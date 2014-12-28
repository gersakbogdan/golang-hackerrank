package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(nr uint64) bool {
	s := math.Sqrt(float64(nr))
	return  uint64(s) * uint64(s) == nr
}

func isFibo(nr uint64) bool {
	return isPerfectSquare(5 * nr * nr + 4) || isPerfectSquare(5 * nr * nr - 4)
}

func isFibo2(nr uint64) bool {
	var a, b uint64
	a, b = 1, 2

	for b < nr {
		a, b = b, a + b

		if b == nr {
			return true
		} else if b > nr {
			return false
		}
	}

	return false
}

// 7778742049
func main() {
	var n, nr uint64
	fmt.Scanf("%v", &n)

	for n > 0 {
		fmt.Scanf("%v", &nr)

		if nr == 1 || nr == 2 || nr == 3 || isFibo2(nr) {
			fmt.Println("IsFibo")
		} else {
			fmt.Println("IsNotFibo")
		}

		n--
	}
}