package main

import (
	"fmt"
)

func main() {
    var N, A, B int
    fmt.Scanf("%d %d %d", &N, &A, &B)
	
	syo := N / (A + B)
	amari := N % (A + B)

	ans := A * syo
	if amari > A {
		ans += A 
	} else {
		ans += amari
	}

	fmt.Printf("%d\n", ans)
}
