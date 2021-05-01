package main

import (
	"fmt"
	"math"
)

func isInteger(x float64) bool {
	return math.Floor(x) == x
}

func main() {

	var N int
	var D, H float64
	fmt.Scanf("%d %f %f", &N, &D, &H)

	var r, R float64
	var d, h float64

	R = 100000000
	for i := 0; i < N; i++ {
		fmt.Scanf("%f %f", &d, &h)
		r = (H - h) / (D - d)

		if R > r {
			R = r
		}
	}
	m := H - R*D

	if m < 0 {
		fmt.Printf("%f\n", 0.0)
	} else {
		fmt.Printf("%f\n", m)
	}

}
