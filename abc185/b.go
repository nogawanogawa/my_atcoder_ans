package main

import (
	"fmt"
)

func main() {
	var N, M int64
	var T float64
    fmt.Scanf("%d %d %f", &N, &M, &T)
	
	var i, n_ int64
	var A, B float64
	var B_ float64 = 0
	var n int64 = N

	var flag bool = true
	for i=0; i<M; i++{
		fmt.Scanf("%f %f", &A, &B)

		// 店に入る前までの処理
		n_ = int64((A - B_) + 0.5)
		n = n - n_

		if n < 1{
			flag = false
			break
		}

		// 店に入った後の処理
		n_ = int64((B - A) + 0.5)
		n = n + n_

		B_ = B

		if n > N{
			n = N
		} 
	}

	n_ = int64((T - B_) + 0.5)
	n = n - n_

	if n < 1{
		flag = false
	}

	if flag { 
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
	
}
