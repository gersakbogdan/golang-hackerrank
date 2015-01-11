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
    2 4
    ---
    2 2

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

    1 5 4 3 2 6
    1 2 3 4 5 6
*/

package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

func sort(arr []int) []int {
    l := len(arr)

    if l <= 1 {
        return arr
    }

    left := make([]int, 0)
    right := make([]int, 0)
    m := ceilDiv(l, 2)

    left = append(left, arr[0:m]...)
    right = append(right, arr[m:l]...)

    nLeft := sort(left)
    nRight := sort(right)

    return merge(nLeft, nRight)
}

func merge(left, right []int) []int {
    li := 0
    ri := 0
    ll := len(left)
    rl := len(right)

    mergeArr := make([]int, 0)

    for li < ll && ri < rl {
        if left[li] < right[ri] {
            mergeArr = append(mergeArr, left[li])
            li++
        } else {
            mergeArr = append(mergeArr, right[ri])
            ri++
        }
    }
    for li < ll {
        mergeArr = append(mergeArr, left[li])
        li++
    }
    for ri < rl {
        mergeArr = append(mergeArr, right[ri])
        ri++
    }

    return mergeArr
}

func ceilDiv(a, b int) int {
    return int(math.Ceil(float64(a) / float64(b)))
}

func almostSorted(sarr, arr []int) string {
    i := 0
    l := len(arr)
    j := l - 1
    br := 0

    for i < j && br < 2 {
        br = 0

        if arr[i] == sarr[i] {
            i++
        } else {
            br++
        }

        if arr[j] == sarr[j] {
            j--
        } else {
            br++
        }
    }

    if i == j {
        return fmt.Sprint("yes")
    }

    ii := i + 1
    ij := j + 1
    times := 0

    br = 0

    for i < j && br < 1 {
        br = 0

        if arr[i] == sarr[j] && arr[j] == sarr[i] {
            times++
            i++
            j--
        } else {
            br++
        }
    }

    if i >= j {
        if times == 1 {
            return fmt.Sprintf("yes\nswap %v %v", ii, ij)
        }
        return fmt.Sprintf("yes\nreverse %v %v", ii, ij)
    }

    br = 0
    for i < j && br < 1 {
        if abs(sarr[i], arr[i]) == abs(sarr[j], arr[j]) {
            i++
            j--
        } else {
            br++
        }
    }

    if i >= j {
        return fmt.Sprintf("yes\nswap %v %v", ii, ij)
    }

    return "no"
}

func abs(x, y int) int {
    if x > y {
        return x - y
    }

    return y - x
}

func main() {
    var n int
    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &n)

    arr := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Fscan(io, &arr[i])
    }

    fmt.Println( almostSorted(sort(arr), arr) )
}