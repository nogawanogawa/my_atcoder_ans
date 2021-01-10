package main

import (
	"fmt"
)

func main() {
	var N int
    fmt.Scanf("%d", &N)

	var sum int64 = 0
	A := make([]int64, N)
	B := make([]int64, N)

	for i:=0; i<N; i++{
		fmt.Scanf("%d", &A[i])
	}

	for i:=0; i<N; i++{
		fmt.Scanf("%d", &B[i])
	}

	for i:=0; i<N; i++{
		sum += A[i] * B[i]
	}

	if sum == 0{
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
	
}
