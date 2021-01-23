package main

import (
	"fmt"
)

func main() {
	var N int
	var X int64
	fmt.Scanf("%d %d", &N, &X)
	X = X*100
	var alcohole int64 = 0
	var v, p int64
	var ans int = -1
	for i:=0; i<N; i++{
		fmt.Scanf("%d %d", &v, &p)
		alcohole += v * p

		if alcohole > X {
			ans = i + 1
			break
		}
	}

	fmt.Printf("%d\n", ans)
}
