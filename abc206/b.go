package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	var sum int = 0
	for i := 1; i <= N; i++ {
		sum = sum + i
		if sum >= N {
			fmt.Printf("%d\n", i)
			break
		}
	}

}
