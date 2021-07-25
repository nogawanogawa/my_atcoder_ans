package main

import (
	"fmt"
)

func main() {

	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	arr := make([][]int, K)
	for i := 0; i < K; i++ {
		arr[i] = make([]int, N)
	}

	C := make([]int, K)
	var c string
	var k int
	for i := 0; i < K; i++ {
		fmt.Scanf("%s %d", &c, &k)
		C[i] = k - 1
		if c == "L" {
			for j := k - 1; j < N; j++ {
				arr[i][j] = 1
			}
		} else {
			for j := k - 1; j >= 0; j-- {
				arr[i][j] = 1
			}
		}
	}

	for i := 0; i < K; i++ {
		k = C[i]

		if arr[i][k] == 0 {
			fmt.Printf("%d\n", 0)
			return
		}

		for j := 0; j < K; j++ {
			arr[j][k] = 0
		}
		arr[i][k] = 1

	}

	var result int = 1
	var sum int = 0

	for i := 0; i < N; i++ {
		sum = 0
		for j := 0; j < K; j++ {
			sum += arr[j][i]
		}

		result = (result * sum) % 998244353
	}
	fmt.Printf("%d\n", result)
}
