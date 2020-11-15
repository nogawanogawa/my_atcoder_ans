package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)
	fmt.Printf("%d", n)

	max := make([]int, n)

	for i:=0; i*i<n; i++{
		if n % i == 0{

			if n/i != i{
				max = append(max, n/i)
			}
			fmt.Printf("%d", i)
		}
	}

	for i:=0; i<n; i++{
		if max[i] != 0{
			fmt.Printf("%d", max[i])
		}
	}
}
