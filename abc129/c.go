package main

import (
	"fmt"
)

func main() {

	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	dp := make([]int, N+1)
	oks := make([]bool, N+1)
	var a int
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &a)
		oks[a] = true
	}

	dp[0] = 1

	for now := 0; now < N; now++ {
		for next := now + 1; (next <= now+2) && (next <= N); next++ {
			if oks[next] == false {
				dp[next] += dp[now]
				dp[next] = dp[next] % 1000000007
			}
		}
	}

	fmt.Printf("%d\n", dp[N])
}
