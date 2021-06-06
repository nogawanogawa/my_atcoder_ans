package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	var A int
	var count int = 0
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A)
		if A > 10 {
			count += A - 10
		}
	}

	fmt.Printf("%d\n", count)

}
