package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

    if n % 2 == 0{
        fmt.Printf("White")
    }else{
        fmt.Printf("Black")
    }
	
}