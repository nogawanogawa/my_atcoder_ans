package main

import (
    "fmt"
)

func main() {

    var N int
    fmt.Scanf("%d", &N)

    if N % 2 == 0{
        fmt.Printf("%f\n", 0.5)
    } else {
        fmt.Printf("%f\n", 1-float64(N/2)/float64(N))
    }

}