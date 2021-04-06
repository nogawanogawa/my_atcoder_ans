package main

import (
    "fmt"
)

func main() {

    var A, B int
    fmt.Scanf("%d %d", &A, &B)

    C := A - 2*B
    if C > 0{
        fmt.Printf("%d\n", C)
    } else {
        fmt.Printf("0\n")
    }

}