package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type man struct {
	index int64
	a     int64
	count int64
}

func main() {

	sc.Split(bufio.ScanWords)

	var N, K int64
	fmt.Scanf("%d %d", &N, &K)

	base := K / N
	K = K - N*base

	A := make([]man, N)
	var i int64

	for i = 0; i < N; i++ {
		A[i].index = i
		A[i].a = nextInt()
		A[i].count = base
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i].a < A[j].a
	})

	for i = 0; i < K; i++ {
		A[i].count += 1
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i].index < A[j].index
	})

	for i = 0; i < N; i++ {
		fmt.Printf("%d\n", A[i].count)
	}

}
