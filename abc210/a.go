package main

import (
	"fmt"
)

func main() {
	var N, A, X, Y int
	fmt.Scanf("%d %d %d %d", &N, &A, &X, &Y)

	var ans int
	if A < N {
		ans = A*X + (N-A)*Y
		fmt.Printf("%d\n", ans)
	} else {
		ans = N * X
		fmt.Printf("%d\n", ans)
	}

}
