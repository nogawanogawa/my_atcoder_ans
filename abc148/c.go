package main

import (
    "fmt"
)

func gcd(a, b int) int {
	if b == 0 {
	  return a
	}
	return gcd(b, a % b)
}

func lcm(a int, b int) int{
    return a * b / gcd(a,b)
}

func main() {
    var A, B int
    fmt.Scanf("%d %d", &A, &B)

    ans := lcm(A, B)

    fmt.Printf("%d\n", ans)
}
