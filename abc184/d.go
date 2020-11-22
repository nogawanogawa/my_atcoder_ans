package main

import (
	"fmt"
)

var dp [101][101][101]float64

// 移動の期待値を返す再帰関数
func f(a int, b int, c int) float64{
	// すでに計算していたら
	if dp[a][b][c] != 0{
		return dp[a][b][c]
	}

	// ゴールしていたら移動の期待値は0
	if (a == 100) || (b == 100) || (c == 100) {
		return 0
	} 

	var ans float64 = 0
	
	ans += (f(a+1, b, c) + 1) * float64(a) / float64(a + b + c)
	ans += (f(a, b+1, c) + 1) * float64(b) / float64(a + b + c)
	ans += (f(a, b, c+1) + 1) * float64(c) / float64(a + b + c)

	dp[a][b][c] = ans
	return ans

}

func main() {

	var A, B, C int
    fmt.Scanf("%d %d %d", &A, &B, &C)

	fmt.Printf("%f\n", f(A, B, C))

}
