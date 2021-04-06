package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    d := make([]int, N)
    for i:=0; i<N; i++{
        fmt.Scanf("%d", &d[i])
    }

    var result int = 0
    for i:=0; i<N-1; i++{
        for j:=i+1; j<N; j++{
            result += d[i] * d[j]
        }     
    }

    fmt.Printf("%d\n", result)
}