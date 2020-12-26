package main

import (
    "fmt"
)

func main() {

    var A, B int
    fmt.Scanf("%d", &A)
    fmt.Scanf("%d", &B)

    var ans int
    if A == 1 {
        if B == 2{
            ans = 3
        } else {
            ans = 2
        }
    } else if A == 2{
        if B ==1 {
            ans = 3
        } else {
            ans = 1
        }
    } else if A == 3{
        if B == 1{
            ans = 2
        } else { 
            ans = 1
        }
    }
    fmt.Printf("%d\n", ans)
}