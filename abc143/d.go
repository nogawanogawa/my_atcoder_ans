package main

import (
    "fmt"
    "sort"
)


func main() {
    var N int
    fmt.Scanf("%d", &N)

    L := make([]int, N)
    for i:=0; i<N; i++{
        fmt.Scanf("%d", &L[i])
    }

    sort.Slice(L, func(i, j int) bool {
        return L[i] > L[j]
    })

    var count int = 0
    for i:=0; i<N-2; i++{
        for j:=i+1; j<N-1; j++{
            if L[i] >= L[j] + L[j+1]{
                break
            }
            for k:=j+1; k<N; k++{
                if L[i] < L[j] + L[k]{
                    count += 1
                } else{
                    break
                }
            }
        }
    }

    fmt.Printf("%d\n", count)  
}
