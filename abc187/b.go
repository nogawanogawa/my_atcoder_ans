package main

import (
	"fmt"
)

func main() {
	var N int
    fmt.Scanf("%d", &N)

	x := make([]int, N)
	y := make([]int, N)

	for i:=0; i<N; i++{
		fmt.Scanf("%d %d", &x[i], &y[i])
	}

	var count int = 0
	for i:=0; i<N-1; i++{
		for j:=i+1; j<N; j++{
			r := float32(y[j] - y[i]) / float32(x[j] - x[i])
			if r * r <= 1 {
				count ++ 
			}
		}
	}

	fmt.Printf("%d\n", count)
	
}
