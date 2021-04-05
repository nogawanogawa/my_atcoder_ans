package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    var flag bool = false
    for i:=1; i<10; i++{
        if N % i == 0{
            if N/i < 10{
                flag = true
            }
        }
    }

    if flag == true{
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }
}