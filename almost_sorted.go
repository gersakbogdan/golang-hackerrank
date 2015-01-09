/*
	Problem Statement

	Given an array with n elements, can you sort this array in ascending order using just one of
	the following operations? You can perform only one of the following operations:
	1. Swap two elements.
	2. Reverse one sub-segment

	Input Format
	The first line contains a single integer n, which indicates the size of the array.
	The next line contains n integers seperated by spaces.

	n
	d1 d2 ... dn

	Constraints

	2 <= n <= 100000
	0 <= di <= 1000000

	All di are distinct.

	Output Format

	    If the array is already sorted, output 'yes' in the first line.
	    You do not need to output anything else.
	    If you can sort this array using one single operation (from the two permitted operations):

	    a. If you can sort the array by swap dl and dr, output "swap l r" in the second line. l and r are the indices of the elements to be swapped, assuming that the array is indexed from 1 to n.
	    b. Else if it is possible to sort the array by reversing the segment d[l...r], output "reverse l r" in the second line. l and r are the indices of the first and last elements of the subsequence to be reversed, assuming that the array is indexed from 1 to n.

	    d[l...r] represents the sub-sequence of the array, beginning at index l and ending at index r; inclusive of both.

	    If an array can be sorted by either swapping or reversing, stick to the swap based method.

	    If you cannot sort the array in either of the above ways, output "no" in the first line.

	Sample Input #1
	2
	4 2

	Sample Output #1
	yes
	swap 1 2


	Sample Input #2
	3
	3 1 2

	Sample Output #2
	no


	Sample Input #3
	6
	1 5 4 3 2 6

	Sample Output #3
	yes
	reverse 2 5


	Explanation

	For #1: You can both swap(1, 2) and reverse(1, 2), but if you can sort the array but swap, output swap only.
	For #2: It is impossible to sort by one single operation (among those permitted).
	For #3, You can reverse the sub-array d[2...5] = "5 4 3 2" then the array become sorted.
*/

package main

import (
	"fmt"
	"bufio"
	"os"
)

func insertionSort(arr []int) {
    l := len(arr)
    swaps := 0
    firstLeftIndex := -1
    leftIndex := -1
    rightIndex := -1
    possible := 1
    swapMethod := 0

    for i := 1; i < l; i++ {
        x := arr[i]
        j := i

        for j > 0 && arr[j - 1] > x {
            arr[j] = arr[j-1]
            j--
        }
        arr[j] = x

        if (j != i) {
        	if leftIndex >= 0 && ((j != leftIndex && j != leftIndex+1 && j != firstLeftIndex) || i != rightIndex + 1) {
        		possible = 0
                break
        	}
            if j == leftIndex + 1 {
                swapMethod = 1
            }
        	leftIndex = j
        	rightIndex = i
            if (firstLeftIndex < 0) {
                firstLeftIndex = j
            }
        	swaps++
            fmt.Println("li:", leftIndex, "ri:", rightIndex, "fli:", firstLeftIndex, "j:", j, "i:", i)
        }
    }

    if possible == 0 {
    	fmt.Println("no")
    } else if swaps == 0 {
    	fmt.Println("yes")
    } else if swaps == 1 {
    	fmt.Println("yes")
    	fmt.Println("swap", leftIndex + 1, rightIndex + 1)
    } else if swaps > 1 && leftIndex != firstLeftIndex {
        fmt.Println("no")
    } else if swaps > 1 && swapMethod == 1 && leftIndex == firstLeftIndex {
        fmt.Println("yes")
        fmt.Println("swap", leftIndex + 1, rightIndex + 1)
    } else {
    	fmt.Println("yes")
    	fmt.Println("reverse", leftIndex + 1, rightIndex + 1)
    }
}

func main() {
    var n int
    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &n)

    arr := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Fscan(io, &arr[i])
    }

    insertionSort(arr)
}
