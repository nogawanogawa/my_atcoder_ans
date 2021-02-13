package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.ParseInt(sc.Text(),10,32)
    if e != nil {
        panic(e)
    }
    return int(i)
}


func main() {
    sc.Split(bufio.ScanWords)

    var T int
    fmt.Scanf("%d", &T)

    var L, R int 
    var count int

    counts := make([]int, T)
    
    for i:=0; i<T; i++{
        count = 0
        L = nextInt()
        R = nextInt()

        if (L == 0) && (R == 0){
            count = 1
        } else {
            d := R - L
            if (d<=R) && (L<=d){
                count = (d-L+1) * (d-L+2) /2
            }
        }
        counts[i] = count    
    }

    for i:=0; i<T; i++{
        fmt.Printf("%d\n", counts[i])
    }    
}
