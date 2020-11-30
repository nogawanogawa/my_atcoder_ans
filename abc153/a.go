package main

import (
    "fmt"
)

func main() {
    var H, A int

    fmt.Scanf("%d %d", &H, &A)

    var ans int
    if H % A == 0{
        ans = H/A
    } else {
        ans = H/A + 1
    }

	fmt.Printf("%d\n", ans)	
}