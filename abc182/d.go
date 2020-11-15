package main

import (
	"fmt"
)


func main() {

    var H, W, N, M int
    fmt.Scanf("%d %d %d %d", &H, &W, &N, &M)

    A := make([]int, N)
    B := make([]int, N)
	for i := 0; i < N; i++ {
        fmt.Scanf("%d %d", &A[i], &B[i])
	}

	C := make([]int, N)
    D := make([]int, N)
	for i := 0; i < M; i++ {
        fmt.Scanf("%d %d", &C[i], &D[i])
	}

	Matrix := make([][]int, H)
	for i := 0; i < n; i++ {
        Matrix[i] = make([]int, W)
	}

	for i:=0; i<N; i++{
		fmt.Scanf("%d", &num)

		total = total + num 

		if max_right < total { 
			max_right = total
		}

		if ans < p_now + max_right {
			ans = p_now + max_right
		}
		
		p_now = p_now + total
	}

	fmt.Printf("%d\n", ans)
}
