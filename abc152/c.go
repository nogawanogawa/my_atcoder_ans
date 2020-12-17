package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    P := make([]int, N)
    P_min := make([]int, N)
    for i:=0; i<N; i++{
        fmt.Scanf("%d", &P[i])
    }

    P_min[0] = P[0]

    for i:=1; i<N; i++{
        if P_min[i-1] > P[i]{
            P_min[i] = P[i]
        } else {
            P_min[i] = P_min[i-1]
        }
    }

    var count int = 1
    for i:=1; i<N; i++{
        if P[i] < P_min[i-1]{
            count ++
        } 
    }

    fmt.Printf("%d\n", count)
    
}