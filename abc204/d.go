package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	T := make([]int, N)
	var sum int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &T[i])
		sum += T[i]
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, sum/2+1)
	}

	for i := 0; i < N; i++ {
		if i == 0 {
			for j := 0; j <= sum/2; j++ {
				if j < T[i] {
					dp[i][j] = 0
				} else {
					dp[i][j] = T[i]
				}
			}
		} else {
			for j := 0; j <= sum/2; j++ {
				if j-T[i] < 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					a := dp[i-1][j]
					b := dp[i-1][j-T[i]] + T[i]
					if a > b {
						dp[i][j] = a
					} else {
						dp[i][j] = b
					}
				}
			}
		}
	}
	max := dp[N-1][sum/2]
	fmt.Printf("%d\n", sum-max)

}
