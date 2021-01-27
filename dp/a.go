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

    var N int
    fmt.Scanf("%d", &N)

    h := make([]int, N)
    for i:=0; i<N; i++{
        h[i] = nextInt()
    }

	dp := make([]int, N)

	dp[0] = 0
	dp[1] = int(math.Abs(float64(h[0] - h[1])))

    for i:=2; i<N; i++{
		dp_1 := int(math.Abs(float64(h[i] - h[i-1]))) + dp[i-1]
		dp_2 := int(math.Abs(float64(h[i] - h[i-2]))) + dp[i-2]

		if dp_1 < dp_2 {
			dp[i] = dp_1
		} else {
			dp[i] = dp_2
		}
	}

	fmt.Printf("%d\n", dp[N-1])
}