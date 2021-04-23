package main

import (
	"fmt"
)

func main() {
	var X, A int

	fmt.Scanf("%d %d", &X, &A)

	if X < A {
		fmt.Printf("0\n")
	} else if X >= A {
		fmt.Printf("10\n")
	}

}
