/*
    Problem Statement

    Sherlock is given an array of N integers A0, A1 ... AN-1 by Watson.
    Now Watson asks Sherlock how many different pairs of indices i and j exist such that i is not equal to j but Ai is equal to Aj.

    That is, Sherlock has to count total number of pairs of indices (i, j) where Ai = Aj AND i ≠ j.

    Input Format
    First line contains T, the number of testcases. T test case follows.
    Each testcase consists of two lines, first line contains an integer N, size of array.
    Next line contains N space separated integers.

    Output Format
    For each testcase, print the required answer in different line.

    Constraints
    1 ≤ T ≤ 10
    1 ≤ N ≤ 105
    1 ≤ A[i] ≤ 106

    Sample input
    2
    3
    1 2 3
    3
    1 1 2

    Sample output
    0
    2

    Explanation
    In the first testcase, no two pair of indices exist which satisfy the given property.
    In second testcase as A[0] = A[1] = 1, the pairs of indices (0,1) and (1,0) satisfy the given property.

    Conclusion: Find duplicates and arrangements
 */
package main

import (
	"fmt"
	"bufio"
	"os"
)

type Counter struct {
	m map[int]int
}

func NewCounter() *Counter {
	return &Counter{make(map[int]int)}
}

func (mp *Counter ) Add(v int) {
	mp.m[v]++
}

func (mp *Counter) Get(v int) int {
	return mp.m[v]
}

func (mp *Counter) Pairs() int {
	total := 0

	for _, v := range mp.m {
		if v > 1 {
			total += v * (v - 1)
		}
	}

	return total
}

func main() {
	var t, n, temp int

	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &t)

	for t > 0 {
		fmt.Fscan(io, &n)
		counter := NewCounter()

		for i := 0; i < n; i++ {
			fmt.Fscan(io, &temp)
			counter.Add(temp)
		}

		fmt.Println(counter.Pairs())

		t--
	}
}
