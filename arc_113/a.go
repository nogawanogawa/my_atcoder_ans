package main

import (
    "fmt"
)

func enum_dividors(AB int) int{

    var num int = 0
    var d int
    for i:=1; i*i<=AB; i++{
        if AB%i == 0{
            num ++
            d = i
        }
    }

    if (AB % d == 0) && (AB / d == d){
        num = num * 2 - 1
    } else {
        num = num * 2
    }


    return num
}

func main() {
    var K int
    fmt.Scanf("%d", &K)

    var num, c_num, ans int
    for AB:=1; AB<=K; AB++{
        num = enum_dividors(AB)
        c_num = K / AB
        ans += num * c_num 
    }

    fmt.Printf("%d\n", ans)
}
