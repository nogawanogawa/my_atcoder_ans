package main

import (
	"fmt"
)

func main() {
	var H, W int
    fmt.Scanf("%d %d", &H, &W)

	var A, sum int
	var min int = 100
	for i:=0; i<H; i++{
		for j:=0; j<W; j++{
			fmt.Scanf("%d", &A)
			if A < min {
				min = A
			}
			sum += A
		}
	}

	ans := sum - min * H * W

	fmt.Printf("%d\n", ans)
	
}
