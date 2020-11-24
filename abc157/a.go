package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    var ans int

    if N % 2 == 0{
        ans = N / 2
    } else {
        ans = N / 2 + 1
    }

	fmt.Printf("%d\n", ans)	
}