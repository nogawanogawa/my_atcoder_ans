package main

import (
    "fmt"
    "strings"
)

func main() {
    var N int
    var S, T string
    fmt.Scanf("%d", &N)
    fmt.Scanf("%s %s", &S, &T)

    s := strings.Split(S,"")
    t := strings.Split(T,"")

    var ans string
    for i:=0; i<N; i++{
        ans += s[i]
        ans += t[i]
    }

    fmt.Printf("%s\n", ans)

}