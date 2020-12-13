package main

import (
    "fmt"
)

func main() {
    var a, b int 
    fmt.Scanf("%d", &a)

    for i:=0; i<3; i++{
        fmt.Scanf("%d", &b)
        if b < a {
            a = b
        }
    }

	fmt.Printf("%d\n", a)
	
}