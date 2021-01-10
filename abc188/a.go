package main

import (
    "fmt"
)

func main() {
    var X, Y int
    fmt.Scanf("%d %d", &X, &Y)

    var flag bool = false
    if X>Y{
        if X-Y < 3{
            flag = true
        }
    } else {
        if Y-X < 3{
            flag = true
        }
    }

    if flag{
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }
}