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
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func main() {

	sc.Split(bufio.ScanWords)

	var N int64
	N = nextInt()

	F := make([]int64, N+2)
	F[0] = 1
	F[1] = 1

	A := make([]int64, N)
	var i int64

	for i = 0; i < N; i++ {
		A[i] = nextInt()
		F[i+2] = (F[i+1] + F[i]) % 1000000007
	}

	var ans int64 = 0
	ans = (ans + F[N]*A[0]) % 1000000007
	var count int64
	for i = 1; i < N; i++ {
		count = ((F[i]*F[N-i])%1000000007 - (F[i-1]*F[N-i-1])%1000000007) % 1000000007
		if count < 0 {
			count += 1000000007
		}
		ans = (ans + (count*A[i])%1000000007) % 1000000007
	}

	fmt.Printf("%d\n", ans)

}
