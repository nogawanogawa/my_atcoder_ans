package main

import (
	"fmt"
)

func main() {

	var N int64
	fmt.Scanf("%d", &N)

	var a, b, c, i int64
	var B int64
	var ans int64 = 1000000000000000000
	for b = 0; b < N; b++ {
		B = 1
		for i = 0; i < b; i++ {
			B *= 2
		}
		if B > N {
			break
		}
		a = N / B
		c = N - a*B
		if ans > a+b+c {
			ans = a + b + c
		}
	}
	fmt.Printf("%d\n", ans)

}
