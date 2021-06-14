package main

import (
	"fmt"
)

func main() {
	var A, B float64
	fmt.Scanf("%f %f", &A, &B)

	ans := B / 100 * A

	fmt.Printf("%f\n", ans)

}
