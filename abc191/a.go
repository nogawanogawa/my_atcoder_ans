package main

import (
    "fmt"
)

func main() {
    var V, T, S, D int
    fmt.Scanf("%d %d %d %d", &V, &T, &S, &D)

    start := V * T
    end := V * S

    if (start > D) || (end < D) {
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }

}