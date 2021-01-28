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

    a := make([]int64, N)
    b := make([]int64, N)
    c := make([]int64, N)
    for i:=0; i<N; i++{
        a[i] = nextInt()
        b[i] = nextInt()
        c[i] = nextInt()
    }

    dp := make([][]int64, N)
    for i:=0; i<N; i++{
        dp[i] = make([]int64, 3)        
    }

    dp[0][0] = a[0]
    dp[0][1] = b[0]
    dp[0][2] = c[0]

    for i:=1; i<N; i++{
        for j:=0; j<3; j++{
            if j == 0{
                if dp[i-1][1] < dp[i-1][2]{
                    dp[i][j] = dp[i-1][2] + a[i]
                } else {
                    dp[i][j] = dp[i-1][1] + a[i]
                }
            } else if j == 1{
                if dp[i-1][0] < dp[i-1][2]{
                    dp[i][j] = dp[i-1][2] + b[i]
                } else {
                    dp[i][j] = dp[i-1][0] + b[i]
                }
            } else {
                if dp[i-1][0] < dp[i-1][1]{
                    dp[i][j] = dp[i-1][1] + c[i]
                } else {
                    dp[i][j] = dp[i-1][0] + c[i]
                }
            }
        }        
    }

    if (dp[N-1][0] > dp[N-1][1]) && (dp[N-1][0] > dp[N-1][2]){
        fmt.Printf("%d\n", dp[N-1][0])
    } else if (dp[N-1][1] > dp[N-1][0]) && (dp[N-1][1] > dp[N-1][2]){
        fmt.Printf("%d\n", dp[N-1][1])
    } else {
        fmt.Printf("%d\n", dp[N-1][2])
    }
}