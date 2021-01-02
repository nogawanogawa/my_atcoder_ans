package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

type toh struct {
    a, t int
}

var r = bufio.NewReader(os.Stdin)

func main() {
    var n int
    fmt.Fscan(r, &n)
    ao := 0
    ar := make([]toh, n)
    for i := range ar {
        var a, b int
        fmt.Fscan(r, &a, &b)
        ao += a
        ar[i].a = a
        ar[i].t = b
    }
    sort.Slice(ar, func(i, j int) bool {
        return ar[i].a+ar[i].t+ar[i].a < ar[j].a+ar[j].t+ar[j].a
    })
    ans := 0
    t := 0
    for i := n - 1; i >= 0; i-- {
        ao -= ar[i].a
        t += ar[i].a + ar[i].t
        ans++
        if ao < t {
            fmt.Println(ans)
            return
        }
    }
}