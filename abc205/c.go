package main

import (
	"fmt"
)

func main() {
	var A, B, C int
	fmt.Scanf("%d %d %d", &A, &B, &C)

	if A == B {
		fmt.Printf("=\n")
	} else if A == -B {
		if C%2 == 0 {
			fmt.Printf("=\n")
		} else {
			if A > B {
				fmt.Printf(">\n")
			} else {
				fmt.Printf("<\n")
			}
		}
	} else if A > 0 {
		if A < B {
			fmt.Printf("<\n")
		} else if (-A < B) && (B < A) {
			fmt.Printf(">\n")
		} else if B < -A {
			if C%2 == 0 {
				fmt.Printf("<\n")
			} else {
				fmt.Printf(">\n")
			}
		}
	} else if A == 0 {
		if B > 0 {
			fmt.Printf("<\n")
		} else {
			if C%2 == 0 {
				fmt.Printf("<\n")
			} else {
				fmt.Printf(">\n")
			}
		}
	} else if A < 0 {
		if -A < B {
			fmt.Printf("<\n")
		} else if (A < B) && (B < -A) {
			if C%2 == 0 {
				fmt.Printf(">\n")
			} else {
				fmt.Printf("<\n")
			}
		} else if B < A {
			if C%2 == 0 {
				fmt.Printf("<\n")
			} else {
				fmt.Printf(">\n")
			}
		}
	}

}
