package main

import (
    "fmt"
)

func main() {
    var S, P int64
    fmt.Scanf("%d %d", &S, &P)

    var ans string = "No"
    var N int64 
    for N=1; N*N<=P; N++{
        M := S - N
        if P == N * M{
            ans = "Yes"
        }
    }
    fmt.Printf("%s\n", ans)
}