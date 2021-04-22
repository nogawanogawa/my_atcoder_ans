package main

import (
	"fmt"
)

func main() {
	var N, L int

	fmt.Scanf("%d %d", &N, &L)

	min := L + 1 - 1
	max := L + N - 1

	var taste int
	if min >= 0 {
		taste = (((L + 2 - 1) + (L + N - 1)) * (N - 1)) / 2
	} else if (min < 0) && (max > 0) {
		taste = (((L + 1 - 1) + (L + N - 1)) * N) / 2
	} else {
		taste = (((L + 1 - 1) + (L + N - 2)) * (N - 1)) / 2
	}

	fmt.Printf("%d\n", taste)

}
