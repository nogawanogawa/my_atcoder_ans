package main

import (
	"fmt"
)

func main() {

	var A1, A2, A3 int
	fmt.Scanf("%d %d %d", &A1, &A2, &A3)

	var ans int
	if A1 < A3 {
		A := A2 - A1
		B := A3 - A2

		if A > B {
			ans = A2 + A - A3
		} else {
			// mid wo ugokas
			if (B-A)%2 == 0 {
				ans = (A3+A1)/2 - A2
			} else {
				ans = (A3+A1+1)/2 - A2 + 1
			}
		}
	} else {
		A := A1 - A2
		B := A2 - A3

		if B > A {
			ans = B - A
		} else {
			// mid wo ugokas
			if (A-B)%2 == 0 {
				ans = (A1+A3)/2 - A2
			} else {
				ans = (A3+A1+1)/2 - A2 + 1
			}
		}
	}

	fmt.Printf("%d\n", ans)

}
