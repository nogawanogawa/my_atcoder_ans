package main

import (
	"fmt"
	"strconv"
)

func main() {

	var N, K int
	fmt.Scanf("%d %d", &N, &K)
	var s string
	for i := 0; i < K; i++ {
		if N%200 == 0 {
			N = N / 200
		} else {
			s = strconv.Itoa(N)
			s = s + "200"
			N, _ = strconv.Atoi(s)
		}
	}

	fmt.Printf("%d\n", N)

}
