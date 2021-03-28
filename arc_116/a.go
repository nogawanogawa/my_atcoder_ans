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

    var T int
    fmt.Scanf("%d", &T)

    N := make([]int64, T)
    for i:=0; i<T; i++{
        N[i] = nextInt()
    }

    for i:=0; i<T; i++{
        if N[i] % 4 == 0{
            fmt.Printf("Even\n")
        } else if N[i] % 2 == 0 {
            fmt.Printf("Same\n")
        } else {
            fmt.Printf("Odd\n")
        }
    }

}
