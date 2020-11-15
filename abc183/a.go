package main

import (
    "fmt"
)

func main() {
    var x int 
    fmt.Scanf("%d", &x)

    var ans int
    if x > 0{
        ans = x
    } else{
        ans = 0
    }
    fmt.Printf("%d\n", ans)
	
}