package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

	var count int = 0
	p := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scanf("%d", &p[i])
		if p[i] != i+1{
			count ++
		}
	}

	if count < 3{
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
}
