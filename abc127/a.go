package main

import (
	"fmt"
)

func main() {
	var A, B int

	fmt.Scanf("%d %d", &A, &B)

	if A < 6 {
		fmt.Printf("%d\n", 0)
	} else if A > 12 {
		fmt.Printf("%d\n", B)
	} else {
		fmt.Printf("%d\n", B/2)
	}

}
