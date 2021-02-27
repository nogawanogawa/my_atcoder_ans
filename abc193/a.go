package main

import (
    "fmt"
)

func main() {
    var A, B float32
    fmt.Scanf("%f %f", &A, &B)

    wariai := (A-B)/A *100

    fmt.Printf("%f\n", wariai)
}