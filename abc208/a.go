package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	min := A * 1
	max := A * 6

	if (min <= B) && (B <= max) {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}

}
