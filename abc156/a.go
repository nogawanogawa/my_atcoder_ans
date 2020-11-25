package main

import (
    "fmt"
)

func main() {
    var N, R int
    fmt.Scanf("%d %d", &N, &R)

    var ans int

	if N >= 10 {
        ans = R
    } else {
        ans = R + (100 * (10 - N))
    }

	fmt.Printf("%d\n", ans)	
}