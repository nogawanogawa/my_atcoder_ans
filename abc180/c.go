package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

	max := make([]int, 1)

	for i:=1; i*i<n+1; i++{
		if n % i == 0{
			if n/i != i{
				max = append(max, n/i)
			}
			fmt.Printf("%d\n", i)
		}
	}


	for i:=len(max)-1; i>0; i--{
		if max[i] != 0{
			fmt.Printf("%d\n", max[i])
		}
	}
}
