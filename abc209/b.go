package main

import (
	"fmt"
)

func main() {

	var N, X int
	fmt.Scanf("%d %d", &N, &X)

	var A int = 0
	var a int

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a)
		if i%2 == 1 {
			A += a - 1
		} else {
			A += a
		}
	}

	if A > X {
		fmt.Printf("No\n")
	} else {
		fmt.Printf("Yes\n")
	}
}
