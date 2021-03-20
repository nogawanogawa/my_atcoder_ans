package main

import (
    "fmt"
)

func main() {
    var a, b, c, d int
    fmt.Scanf("%d %d", &a, &b)
    fmt.Scanf("%d %d", &c, &d)

    e := a - c
    f := a - d
    g := b - c
    h := b - d

    if (e >= f) && (e >= g) && (e >= h){
        fmt.Printf("%d\n", e)
    } else if (f >= e) && (f >= g) && (f >= h){
        fmt.Printf("%d\n", f)
    } else if (g >= f) && (g >= e) && (g >= h){
        fmt.Printf("%d\n", g)
    } else {
        fmt.Printf("%d\n", h)
    }

}