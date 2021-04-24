package main

import (
	"fmt"
)

func main() {
	var A, B, C int
	fmt.Scanf("%d %d %d", &A, &B, &C)

	if A*A+B*B < C*C {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")

	}

}
