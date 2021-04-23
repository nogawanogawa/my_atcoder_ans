package main

import (
	"fmt"
)

func main() {
	var N, X int

	fmt.Scanf("%d %d", &N, &X)

	var L int
	var D int = 0
	var count int = 1
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &L)
		if D+L <= X {
			count++
		}
		D = D + L
	}

	fmt.Printf("%d\n", count)

}
