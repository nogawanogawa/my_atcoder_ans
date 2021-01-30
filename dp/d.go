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

    var N, W int64
    fmt.Scanf("%d %d", &N, &W)

    w := make([]int64, N)
    v := make([]int64, N)

    var i,j int64
    for i=0; i<N; i++{
        w[i] = nextInt()
        v[i] = nextInt()
    }

    dp := make([][][]int64, N)
    for i=0; i<N; i++{
        dp[i] = make([][]int64, W+1)
        for j=0; j<=W; j++{
            dp[i][j] = make([]int64, 2)
        }
    }

    for j=0; j<=W; j++{
        if w[0] <= j{
            dp[0][j][0] = w[0]
            dp[0][j][1] = v[0]
        }
    }

    for i=1; i<N; i++{
        for j=0; j<=W; j++{
            if j - w[i] >= 0{
                if dp[i-1][j-w[i]][1] + v[i] >= dp[i-1][j][1]{
                    dp[i][j][0] = dp[i-1][j-w[i]][0] + w[i]
                    dp[i][j][1] = dp[i-1][j-w[i]][1] + v[i]
                } else {
                    dp[i][j][0] = dp[i-1][j][0]
                    dp[i][j][1] = dp[i-1][j][1]
                }
            } else {
                dp[i][j][0] = dp[i-1][j][0]
                dp[i][j][1] = dp[i-1][j][1]
            }
        }        
    }

    fmt.Printf("%d\n", dp[N-1][W][1])
}