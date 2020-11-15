package main

import (
    "fmt"
)

func main() {
    var n int64
    fmt.Scanf("%d", &n)

    var count_a int64 = 0
    var count_b int64 = 0

    var a, b int64 
    flag := true
    for a=3; a<=n; a=a*3{
        count_a += 1
        count_b = 0
        for b=5; b<n-a+1; b=b*5{
            count_b += 1
            if (a+b == n) && flag{
                flag = false
                fmt.Printf("%d %d\n", count_a, count_b)
                break
            }    
        }
    }

    if flag{
        fmt.Printf("%d\n", -1)
    }
}