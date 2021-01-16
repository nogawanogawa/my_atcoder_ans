package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int64 {
    sc.Scan()
    i, e := strconv.ParseInt(sc.Text(),10,64)
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
    sc.Split(bufio.ScanWords)

    var N int
    fmt.Scanf("%d", &N)

    var a int64
    var c int64
    b := make([]int64, N)
    d := make([]int64, N)

    a = nextInt()
    d[0] = a
    for i:=1; i<N; i++ {
        a = nextInt()
        if d[i-1] > a{
            d[i] = d[i-1]
        } else {
            d[i] = a
        }
    }

    for i:=0; i<N; i++ {
        b[i] = nextInt()
    }

    c = d[0] * b[0]
    fmt.Printf("%d\n", c)

    for i:=1; i<N; i++{

        if d[i] * b[i] > c{
            c = d[i] * b[i]
        } else {
            c = c
        }
        fmt.Printf("%d\n", c)
    }
}
