package main

import (
	"fmt"
)

func main() {
	var N, A, B int

	fmt.Scanf("%d %d %d", &N, &A, &B)

	if A*N > B {
		fmt.Printf("%d\n", B)
	} else {
		fmt.Printf("%d\n", A*N)
	}
}
