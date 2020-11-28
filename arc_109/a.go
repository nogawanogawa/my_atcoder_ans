package main

import (
    "fmt"
)

func main() {
    var a,b,x,y int
    fmt.Scanf("%d %d %d %d", &a, &b, &x, &y)

    var ans int 
    if a == b {
        ans = x
    } else {
        if a > b{
            ans = (a - (b+1)) * y + x    
            if ans > (2 * (a-b) -1) * x {
                ans = (2 * (a-b) -1) * x
            }    
        } else {
            ans = (b - a) * y + x
            if ans > ((1 + 2 * (b-a)) * x) {
                ans = ((1 + 2 * (b-a)) * x)
            }
        }    
    }

    fmt.Printf("%d\n", ans)
}