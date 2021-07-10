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

	C := make([]int, N)

	for i := 0; i < N; i++ {
		C[i] = nextInt()
	}

	sort.Slice(C, func(i, j int) bool {
		return C[i] < C[j]
	})

	var result int = 1

	for i := 0; i < N; i++ {
		result = result * (C[i] - i) % 1000000007
	}

	fmt.Printf("%d\n", result)

}
