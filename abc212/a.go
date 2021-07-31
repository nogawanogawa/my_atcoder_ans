package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	if (0 < A) && (B == 0) {
		fmt.Printf("Gold\n")
	} else if (0 == A) && (0 < B) {
		fmt.Printf("Silver\n")
	} else {
		fmt.Printf("Alloy\n")
	}

}
