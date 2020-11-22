package main

import (
    "fmt"
)

func main() {
    var A, B int
    fmt.Scanf("%d %d", &A, &B)

	var n int = 1
	var ans int
	if B  == 1 {
		ans = 0
	} else {
		for i := 1; i<B ; i++{
			n = n - 1 + A
			if n >= B {
				ans = i
				break
			}
		}
	}

    fmt.Printf("%d\n", ans)
}
