package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

	var A float64
	var mul float64 = 0
	for i:=0; i<N; i++{
		fmt.Scanf("%f", &A)
		mul += 1/A
	}

	ans := 1/mul

    fmt.Printf("%f\n", ans)
}
