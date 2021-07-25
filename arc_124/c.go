package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	var N int
	fmt.Scanf("%d", &N)

	a := make([]int, N)
	b := make([]int, N)

	var i int

	for i = 0; i < N; i++ {
		fmt.Scanf("%d %d", &a[i], &b[i])
	}

	var l_a, l_b []int

	A := a[0]
	B := b[0]

	l_a = append(l_a, A)
	l_b = append(l_b, B)

	for i = 1; i*i <= A; i++ {
		if A%i == 0 {
			l_a = append(l_a, i)
			if i*i != A {
				l_a = append(l_a, A/i)
			}
		}
	}

	for i = 1; i*i <= B; i++ {
		if B%i == 0 {
			l_b = append(l_b, i)
			if i*i != B {
				l_b = append(l_b, B/i)
			}
		}
	}

	var ans int
	for _, d1 := range l_a {
		for _, d2 := range l_b {

			ok := true
			for j := 1; j < N; j++ {
				a_, b_ := a[j], b[j]

				if (a_%d1 == 0 && b_%d2 == 0) || (b_%d1 == 0 && a_%d2 == 0) {
					continue
				}
				ok = false
				break
			}
			if ok {
				ans = max(ans, lcm(d1, d2))
			}
		}
	}

	fmt.Printf("%d\n", ans)

}
