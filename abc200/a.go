package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	if N%100 == 0 {
		fmt.Printf("%d\n", N/100)
	} else {
		fmt.Printf("%d\n", (N/100)+1)

	}

}
