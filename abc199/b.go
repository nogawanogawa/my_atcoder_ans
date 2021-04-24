package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	A := make([]int, N)
	B := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &B[i])
	}

	var max int = 1000
	var min int = 0

	for i := 0; i < N; i++ {
		if A[i] > min {
			min = A[i]
		}
		if B[i] < max {
			max = B[i]
		}
	}

	if max >= min {
		fmt.Printf("%d\n", max-min+1)
	} else {
		fmt.Printf("%d\n", 0)
	}

}
