package main

import (
	"fmt"
)

func main() {

	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	var ans int = 0
	for i := 1; i < N+1; i++ {
		for j := 1; j < K+1; j++ {
			ans += 100*i + j
		}
	}

	fmt.Printf("%d\n", ans)

}
