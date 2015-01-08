/*
	Problem Statement

	Insertion Sort is a simple sorting technique which was covered in previous challenges.
	Sometimes, arrays may be too large for us to wait around for insertion sort to finish. 
	Is there some other way we can calculate the number of times Insertion Sort shifts each elements 
	when sorting an array?

	If ki is the number of elements over which ith element of the array has to shift then total 
	number of shift will be k1 + k2 + ... + kN.

	Input:
	The first line contains the number of test cases T. T test cases follow. 
	The first line for each case contains N, the number of elements to be sorted. 
	The next line contains N integers a[1],a[2]...,a[N].

	Output:
	Output T lines, containing the required answer for each test case.

	Constraints:
	1 <= T <= 5
	1 <= N <= 100000
	1 <= a[i] <= 1000000

	Sample Input:

	2  
	5  
	1 1 1 2 2  
	5  
	2 1 3 1 2

	Sample Output:

	0  
	4   

	Explanation
	First test case is already sorted, therefore there's no need to shift any element. 
	In second case, it will proceed in the following way.

	Array: 2 1 3 1 2 -> 1 2 3 1 2 -> 1 1 2 3 2 -> 1 1 2 2 3
	Moves:   -        1       -    2         -  1            = 4
*/

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var t, n int

	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &t)

	for t > 0 {
		fmt.Fscan(io, &n)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(io, &arr[i])
		}

		fmt.Println(arr)
		
		t--
	}
}