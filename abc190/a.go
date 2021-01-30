package main

import (
    "fmt"
)

func main() {
    var A, B, C int
    fmt.Scanf("%d %d %d", &A, &B, &C)

    var ans string
    if C == 0{
        max := A
        ans = "Aoki"
        if A == 0{
            ans = "Aoki"
        } else {
            for i:=0; i<max; i++{
                A--
                B--
                if B < 0{
                    ans = "Takahashi"
                    break
                }
            }    
        }
    } else {
        max := B
        ans = "Takahashi"
        if B == 0{
            ans = "Takahashi"
        } else {
            for i:=0; i<max; i++{
                B--
                A--
                if A < 0{
                    ans = "Aoki"
                    break
                }
            }    
        }
    }
    fmt.Printf("%s\n", ans)
}