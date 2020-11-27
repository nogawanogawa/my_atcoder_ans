package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    dict := make(map[int]int)

    var A int
    var flag bool

    for i:=0; i<N; i++{

        fmt.Scanf("%d", &A)
        _, ok := dict[A]
        if ok {
            flag = true
        } else {
            dict[A] = 1
        }
    }

    if flag{
        fmt.Printf("NO\n")
    }else {
        fmt.Printf("YES\n")
    }
    
}