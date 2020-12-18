package main

import (
    "fmt"
)

func main() {

    var C string
    fmt.Scanf("%s", &C)

    for i := 0; i < len(C); i++ {
        a := C[i] + 1
        fmt.Printf("%v\n", string(a))
    }   
}