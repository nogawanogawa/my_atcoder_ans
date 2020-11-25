package main

import (
    "fmt"
)

func main() {
    var N, K int
    fmt.Scanf("%d %d", &N, &K)

    var i int

    for i=1; ; i++{
        D := N / K
        if D == 0{
            break
        }
        N = D
    }

    fmt.Printf("%d\n", i)
}