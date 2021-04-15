package main

import (
	"fmt"
)

func main() {
	var A, B int

	fmt.Scanf("%d %d", &A, &B)

	K := (A + B) / 2
	D := (A + B) % 2

	if D == 0 {
		fmt.Printf("%d\n", K)
	} else {
		fmt.Printf("IMPOSSIBLE\n")
	}
}
