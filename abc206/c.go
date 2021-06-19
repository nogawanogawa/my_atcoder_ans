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
	dict := make(map[int]int)

	for i := 0; i < N; i++ {
		A[i] = nextInt()

		_, ok := dict[A[i]]
		if ok {
			dict[A[i]] = dict[A[i]] + 1
		} else {
			dict[A[i]] = 1
		}
	}

	var sum int = 0
	for i := 0; i < N; i++ {
		a := A[i]
		sum = sum + N - i - dict[a]

		dict[a] = dict[a] - 1
	}

	fmt.Printf("%d\n", sum)
}
