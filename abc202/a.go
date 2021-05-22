package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	A := 7 - a
	B := 7 - b
	C := 7 - c

	fmt.Printf("%d\n", A+B+C)

}
