package main

import (
	"fmt"
)

func main() {
	var A, B, C int

	fmt.Scanf("%d %d %d", &A, &B, &C)

	aki := A-B

	if aki > C {
		fmt.Printf("0\n")
	} else {
		fmt.Printf("%d\n", C-aki)
	}
}
