package main

import (
	"fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)
	
	A := make([]int, N)
	var max int = 2
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		if max < A[i]{
			max = A[i]
		}
	}

	var count, gcd, ans int 

	for i := 2; i<= max; i++ {
		count = 0
		for j:=0; j<N; j++{
			if A[j] % i == 0{
				count +=1
			}
		}

		if count >= gcd {
			ans = i
			gcd = count
		}
	}

	fmt.Printf("%d\n", ans)
	
}
