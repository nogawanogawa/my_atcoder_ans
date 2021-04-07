package main

import (
    "fmt"
    "sort"
)

type student struct{
    id int
    num int
}

func main() {
    var N int
    fmt.Scanf("%d", &N)

    A := make([]student, N)

    for i:=0; i<N; i++{
        fmt.Scanf("%d", &A[i].num)
        A[i].id = i+1
    }

    sort.Slice(A, func(i, j int) bool {
        return A[i].num < A[j].num
    })


    for i:=0; i<N; i++{
        if i!= N-1{
            fmt.Printf("%d ", A[i].id)
        } else {
            fmt.Printf("%d\n", A[i].id)
        }
    }
}
