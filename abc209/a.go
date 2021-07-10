package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	C := B - A + 1

	if C < 0 {
		fmt.Printf("0\n")
	} else {
		fmt.Printf("%d\n", C)
	}

}
