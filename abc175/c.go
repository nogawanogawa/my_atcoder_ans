package main

import (
	"fmt"
)

func main(){
	var X, K, D int64
	fmt.Scanf("%d %d %d",&X, &K, &D)

	var n int64 
	if X > 0 {
		n = X / D
	} else {
		n = (-X) / D
	}


	var ans int64 
	if K <= n{
		diff := K * D
		if X > 0{
			ans = X - diff
		} else {
			ans = -(X + diff)
		}
	} else {
		diff := D * n
		if X > 0{
			X = X - diff
		} else {
			X = X + diff
		}

		if (K-n) % 2 == 0{
			if X > 0{
				ans = X
			} else {
				ans = -X
			}	
		} else {
			if X > 0{
				ans = -(X - D)
			} else {
				ans = X + D
			}	
		}
	} 
	fmt.Printf("%d\n", n)

	fmt.Printf("%d\n", ans)
}
