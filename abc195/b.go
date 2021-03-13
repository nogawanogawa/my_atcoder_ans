package main

import (
    "fmt"
)

func main() {
	
	var A, B, W int
	fmt.Scanf("%d %d %d",&A, &B, &W)

	W = W * 1000
	var min_n int = 10000000
	var max_n int = 0
	for i:=0; i*A<=W; i++{
		if (i * A <= W) && (i * B >= W){
			if min_n > i{
				min_n = i
			}
			if max_n < i{
				max_n = i
			}
		}
	}

	if (min_n == 10000000) || (max_n == 0){
		fmt.Printf("UNSATISFIABLE\n")
	}else {
		fmt.Printf("%d %d\n", min_n, max_n)
	}
}
