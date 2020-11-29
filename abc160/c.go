package main

import (
	"fmt"
)

func main(){
	var K, N int
	fmt.Scanf("%d %d",&K, &N)

	A := make([]int, N)
	B := make([]int, N)

	var max int = 0
	fmt.Scanf("%d",&A[0])
	for i:=1; i<N; i++{
		fmt.Scanf("%d",&A[i])
		B[i-1] = A[i] - A[i-1]
		if max < B[i-1] {
			max = B[i-1]
		}
	}

	B[N-1] = A[0] + (K - A[N-1])
	if max < B[N-1] {
		max = B[N-1]
	}

	fmt.Printf("%d\n", K-max)
}
