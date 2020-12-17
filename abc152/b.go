package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    var a, b int
    fmt.Scanf("%d %d", &a, &b)

    var ans string = "aa"

    A := strconv.Itoa(a)
    B := strconv.Itoa(b)

    if a < b {
        ans = strings.Repeat(A, b)
    } else {
        ans = strings.Repeat(B, a)
    }
    fmt.Printf("%s\n", ans)

}