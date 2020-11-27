package main

import (
    "fmt"
)

func main() {
    var S, T, U string
    var A, B int

    fmt.Scanf("%s %s", &S, &T)
    fmt.Scanf("%d %d", &A, &B)
    fmt.Scanf("%s", &U)

    if S == U{
        A = A-1
    } else if T == U{
        B = B-1
    }

	fmt.Printf("%d %d\n", A, B)	
}