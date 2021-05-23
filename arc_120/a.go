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
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return int(i)
}

func main() {

	sc.Split(bufio.ScanWords)

	var N int
	fmt.Scanf("%d", &N)

	A := make([]int, N)
	var max int = 0
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	var ans int64 = 0
	var tmp int64 = 0

	for i := 0; i < N; i++ {
		if A[i] > max {
			tmp = tmp + int64(A[i]+(A[i]-max))
			ans = ans + tmp + int64(i*(A[i]-max))
			max = A[i]
		} else {
			tmp = tmp + int64(A[i])
			ans = ans + tmp
		}
		fmt.Printf("%d\n", ans)
	}

}
