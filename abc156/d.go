package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	mod := int(1e9) + 7

	ans := (powMod(2, n, mod) - 1 - nCrMod(n, a, mod) - nCrMod(n, b, mod)) % mod
	if ans < 0 {
		ans += mod
	}
	fmt.Println(ans)
}

func powMod(a int, n int, b int) int {
	if n == 1 {
		return a
	}
	if n%2 == 1 {
		return a * powMod(a, n-1, b) % b
	}
	ret := powMod(a, n/2, b)
	ret = (ret * ret) % b

	return ret
}

func nCrMod(n int, r int, mod int) int {
	ret := 1
	//X
	for i := n - r + 1; i <= n; i++ {
		ret *= i
		ret %= mod
	}
	//Y
	for i := 2; i <= r; i++ {
		ret *= powMod(i, mod-2, mod)
		ret %= mod
	}
	ret %= mod
	if ret < 0 {
		ret += mod
	}
	return ret
}
