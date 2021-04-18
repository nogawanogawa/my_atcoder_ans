package main

import (
    "fmt"  
)



func main() {

    var A, B int
    fmt.Scanf("%d %d", &A, &B)

    E_A := make([]int, A)
    E_B := make([]int, B)

    var count int =0
    if A == B {
        for i:=0; i<A; i++{
            E_A[i] = i+1
        }
        for i:=0; i<B; i++{
            E_B[i] = -i-1
        }
    } else if A > B{
        
        for i:=0; i<A; i++{
            E_A[i] = i+1
            if i >= B-1{
                count += i+1
            }
        }
        for i:=0; i<B-1; i++{
            E_B[i] = -i-1
        }
        E_B[B-1] = -count
    } else {
        
        for i:=0; i<A; i++{
            E_A[i] = i+1
        }
        for i:=0; i<B; i++{
            E_B[i] = -i-1
            if i >= A-1{
                count += -i-1
            }
        }
        E_A[A-1] = -count
    }

    for i:=0; i<A; i++{
        fmt.Printf("%d ", E_A[i])
    }
    for i:=0; i<B; i++{
        fmt.Printf("%d ", E_B[i])
    }

}
