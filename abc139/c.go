package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scanf("%d", &N)

	H := make([]int64, N)

	fmt.Scanf("%d", &H[0])

	var max int = 0
	var count int = 0

	for i:=1; i<N; i++{
		fmt.Scanf("%d", &H[i])
		if H[i-1] >= H[i]{
			count += 1
		} else {
			if max < count{
				max = count
			}
			count = 0
		}
	} 

	if max < count{
		max = count
	}

	fmt.Printf("%d\n", max)
}
