package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scanf("%d", &N)

	W := make([]int, N)
	var sum int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &W[i])
		sum += W[i]
	}

	var left, right, diff_ int
	var diff int = 100000
	for i := 0; i < N; i++ {
		left += W[i]
		right = sum - left
		if right > left {
			diff_ = right - left
		} else {
			diff_ = left - right
		}
		if diff > diff_ {
			diff = diff_
		} else {
			break
		}
	}

	fmt.Printf("%d\n", diff)

}
