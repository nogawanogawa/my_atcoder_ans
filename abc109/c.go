package main

import (
	"fmt"
)

func main(){
	var N, X, x int
	fmt.Scanf("%d %d",&N, &X)

	ans := 0
	for i:=0; i<N; i++{
		fmt.Scan(&x)
		if x > X{
			x = x-X
		}else{
			x = X-x
		}

		ans = gcd(ans, x)
	}

	fmt.Printf("%d\n", ans)
}

func gcd(a, b int) int {
	if b == 0 {
	  return a
	}
	return gcd(b, a % b)
}