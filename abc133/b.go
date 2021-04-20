package main

import (
	"fmt"
	"math"
)

func main() {
	var N, D int
	fmt.Scanf("%d %d", &N, &D)

	X := make([][]int, N)
	for i := 0; i < N; i++ {
		X[i] = make([]int, D)
	}

	var count int = 0

	for i := 0; i < N; i++ {
		for j := 0; j < D; j++ {
			fmt.Scanf("%d", &X[i][j])
		}
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			var d int = 0
			for k := 0; k < D; k++ {
				d += (X[i][k] - X[j][k]) * (X[i][k] - X[j][k])
			}

			d_float := math.Sqrt(float64(d))
			if math.Floor(d_float) == d_float {
				count++
			}
		}
	}
	fmt.Printf("%d\n", count)
}
