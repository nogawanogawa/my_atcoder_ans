package main

import (
	"fmt"
)

func main() {
	var A, P int

	fmt.Scanf("%d %d", &A, &P)

	P += A * 3
	pi := P / 2
	fmt.Printf("%d\n", pi)
}
