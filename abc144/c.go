package main

import (
    "fmt"
)


func main() {
    var N int64
    fmt.Scanf("%d", &N)

    var min int64 = 9000000000000
    var i int64
    for i=1; i*i<=N; i++{
        if N % i == 0{
            j:= N/i
            if min > i + j -2{
                min = i+j-2
            }
        }
    }

    fmt.Printf("%d\n", min)
}
