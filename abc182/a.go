package main

import (
    "fmt"
)

func main() {
    var A, B int
    fmt.Scanf("%d %d", &A, &B)

    ans := 2 * A + 100 - B
    fmt.Printf("%d\n", ans)
	
}