package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	p := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &p[i])
	}

	var count int = 0
	for i := 0; i < N-2; i++ {
		if ((p[i] < p[i+1]) && (p[i+1] < p[i+2])) || ((p[i] > p[i+1]) && (p[i+1] > p[i+2])) {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
