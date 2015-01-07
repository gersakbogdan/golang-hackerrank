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