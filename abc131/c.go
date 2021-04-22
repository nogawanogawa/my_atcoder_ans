package main

import (
	"fmt"
)

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int64, b int64) int64 {
	return a * b / gcd(a, b)
}

func main() {

	var A, B, C, D int64
	fmt.Scanf("%d %d %d %d", &A, &B, &C, &D)

	var count int64

	var min, max int64
	if A%C == 0 {
		min = A / C
	} else {
		min = A/C + 1
	}

	max = B / C

	if max != 0 {
		count += max - min + 1
	}

	if A%D == 0 {
		min = A / D
	} else {
		min = A/D + 1
	}

	max = B / D

	if max != 0 {
		count += max - min + 1
	}

	CD := lcm(C, D)

	if A%CD == 0 {
		min = A / CD
	} else {
		min = A/CD + 1
	}

	max = B / CD

	if max != 0 {
		count -= max - min + 1
	}

	fmt.Printf("%d\n", B-A+1-count)
}
