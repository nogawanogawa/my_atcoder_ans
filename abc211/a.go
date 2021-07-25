package main

import (
	"fmt"
)

func main() {
	var A, B float64
	fmt.Scanf("%f %f", &A, &B)

	var ans float64
	ans = (A-B)/3 + B
	fmt.Printf("%f\n", ans)

}
