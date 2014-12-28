package main

import (
	"fmt"
	"sort"
)

type uint64Arr []uint64
func (a uint64Arr) Len() int { return len(a)}
func (a uint64Arr) Swap(i, j int) { a[i], a[j] = a[j], a[i]}
func (a uint64Arr) Less(i, j int) bool { return a[i] < a[j] }


func main() {
	var n, k int
	var arr []uint64

	// get n and k
	fmt.Scan(&n, &k)

	// set an array with length n
	arr = make([]uint64, n)

	// get array values
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	// sort array
	sort.Sort(uint64Arr(arr))

	diff := arr[k-1] - arr[0]

	// find min diff
	for i := 0; i < n - k; i++ {
		cdiff := arr[i + k - 1] - arr[i]
		if cdiff < diff {
			diff = cdiff
		}
	}

	fmt.Println(diff)
}