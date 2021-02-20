package main

import (
    "fmt"
)

func main() {
    var X int
    fmt.Scanf("%d", &X)

    amari := X % 100

    ans := 100 - amari

    fmt.Printf("%d\n", ans)
}