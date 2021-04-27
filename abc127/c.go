package main

import (
	"fmt"
)

func main() {

	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	var L, R int
	var l int = 0
	var r int = 100000
	for i := 0; i < M; i++ {
		fmt.Scanf("%d %d", &L, &R)
		if l < L {
			l = L
		}
		if R < r {
			r = R
		}
	}

	if l <= r {
		fmt.Printf("%d\n", r-l+1)
	} else {
		fmt.Printf("%d\n", 0)
	}

}
