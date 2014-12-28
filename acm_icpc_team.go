package main

import (
	"fmt"
)

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	p := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &p[i])
	}

	topics := 0
	teams := 0
	one := uint8('1')

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ct := 0
			for k := 0; k < m; k++ {
				if p[i][k] == one || p[j][k] == one {
					ct++
				}
			}

			if ct > topics {
				topics = ct
				teams = 1
			} else if ct > 0 && topics == ct {
				teams++
			}
		}
	}

	fmt.Printf("%v\n%v\n", topics, teams)
}