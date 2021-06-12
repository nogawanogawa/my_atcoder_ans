package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	h := N / 2
	x := float64(A[h]) / 2

	var ans float64 = 0
	for i := 0; i < N/2; i++ {
		ans += x
	}

	for i := N / 2; i < N; i++ {
		ans += -x + float64(A[i])
	}

	ans = ans / float64(N)
	fmt.Printf("%f\n", ans)

}
