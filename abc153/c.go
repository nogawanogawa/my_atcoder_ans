package main

import (
    "fmt"
    "sort"
)

func main() {
    var N, K int64
    fmt.Scanf("%d %d", &N, &K)

    H := make([]int64, N)
    
    var i int64
    var sum int64 = 0
    for i=0; i<N; i++{
        fmt.Scanf("%d", &H[i])
        sum += H[i]
    }

    sort.Slice(H, func(i, j int) bool {
        return H[i] > H[j]
    })

    if K > N {
        K = N
    }
    for i=0; i<K; i++{
        sum -= H[i]
    }

    fmt.Printf("%d\n", sum)
    
}