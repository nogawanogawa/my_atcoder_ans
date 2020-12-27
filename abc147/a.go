package main

import (
    "fmt"
)

func main() {

    var A1, A2, A3  int
    fmt.Scanf("%d %d %d", &A1, &A2, &A3)

    if A1+A2+A3 >= 22{
        fmt.Printf("bust\n")
    } else { 
        fmt.Printf("win\n")
    }
}