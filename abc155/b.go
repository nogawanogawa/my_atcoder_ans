package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    var A int
    var flag bool = false
    for i:=0; i<N; i++{
        fmt.Scanf("%d", &A)

        if A%2==0{
            if !(A%3==0) && !(A%5==0){
                flag = true
            }
        }
    }

    if flag {
        fmt.Printf("DENIED\n")
    } else {
        fmt.Printf("APPROVED\n")
    }

}