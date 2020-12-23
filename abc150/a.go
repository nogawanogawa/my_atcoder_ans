package main

import (
    "fmt"
)

func main() {

    var K, X int
    fmt.Scanf("%d %d", &K, &X)

    if K * 500 >= X{
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }
}