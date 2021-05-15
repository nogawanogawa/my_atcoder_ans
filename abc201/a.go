package main

import (
	"fmt"
	"math"
)

func main() {
	var A1, A2, A3 float64
	fmt.Scanf("%f %f %f", &A1, &A2, &A3)

	d1 := math.Abs(A1 - A2)
	d2 := math.Abs(A2 - A3)
	d3 := math.Abs(A3 - A1)

	if (A1 == A2) && (A2 == A3) {
		fmt.Printf("Yes\n")
	} else if (A1 == A2) || (A2 == A3) || (A3 == A1) {
		fmt.Printf("No\n")
	} else if (d1 == d2) || (d2 == d3) || (d3 == d1) {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}

}
