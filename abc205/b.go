package main

import (
	"fmt"
	"sort"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	A := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	for i := 0; i < N; i++ {
		if A[i] != i+1 {
			fmt.Printf("No\n")
			return
		}
	}
	fmt.Printf("Yes\n")

}
