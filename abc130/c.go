package main

import (
	"fmt"
)

func main() {

	var W, H float64
	var x, y int
	fmt.Scanf("%f %f %d %d", &W, &H, &x, &y)

	square := W * H / 2
	var points int = 0

	if (W/2 == float64(x)) && (H/2 == float64(y)) {
		points = 1
	}

	fmt.Printf("%f %d\n", square, points)
}
