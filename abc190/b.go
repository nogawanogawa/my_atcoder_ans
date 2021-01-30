package main

import (
	"fmt"
)

func main() {
	var N int
	var S, D int64

	fmt.Scanf("%d %d %d", &N, &S, &D)

	X := make([]int64, N)
	Y := make([]int64, N)

	var ans string = "No"
	for i:=0; i<N; i++{
		fmt.Scanf("%d %d", &X[i], &Y[i])
		if (X[i] < S) && (Y[i] > D){
			ans = "Yes"
		}
	}

	fmt.Printf("%s\n", ans)
}
