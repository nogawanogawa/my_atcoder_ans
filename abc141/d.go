package main

import (
    "fmt"
    "sort"
)



func main() {
    var N, M int
    fmt.Scanf("%d %d", &N, &M)

    A := make([]int, N)
    for i:=0; i<N; i++{
        fmt.Scanf("%d", &A[i])
    }

    sort.Slice(A, func(i, j int) bool {
        return A[i] > A[j]
    })

    for i:=0; i<M; i++{
        head := A[0]/2
        A = A[1:]
        pos := sort.Search(len(A), func(j int) bool { return A[j] < head})
        if len(A) == 0{
            A = append(A, head)
        } else {
            if pos == 0{
                A, A[0] = append(A[:1], A[0:]...), head
            } else if pos == len(A){
                A = append(A, head)
            } else {
                A = append(A[:pos+1], A[pos:]...)
                A[pos] = head
            } 
        }      
    }
    
    var result int = 0
    for i:=0; i<N; i++{
        result += A[i]
    }

    fmt.Printf("%d\n", result)  
}
