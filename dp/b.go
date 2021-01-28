package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"math"
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

    var N, K int
    fmt.Scanf("%d %d", &N, &K)

    h := make([]int, N)
    for i:=0; i<N; i++{
        h[i] = nextInt()
    }

	dp := make([]int, N)
    for i:=1; i<N; i++{
        dp[i] = -1
    }

	dp[0] = 0

    for i:=1; i<N; i++{
		for j:=1; j<=K; j++{
			if i - j >= 0{
				tmp_dp := int(math.Abs(float64(h[i] - h[i-j]))) + dp[i-j]
	
				if dp[i] == -1{
					dp[i] = tmp_dp
				} else if tmp_dp < dp[i] {
					dp[i] = tmp_dp
				}
			}
		}
	}

	fmt.Printf("%d\n", dp[N-1])
}