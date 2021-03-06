package main

import (
    "fmt"
)

func main() {
    var A, B int
    fmt.Scanf("%d %d", &A, &B)

    C := A + B 

    if (C >= 15) && (B >= 8) {
        fmt.Printf("1\n")
    } else if (C >= 10) && (B >= 3) {
        fmt.Printf("2\n")
    } else if (C >= 3) {
        fmt.Printf("3\n")
    } else {
        fmt.Printf("4\n")
    }

}