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

	A := make(map[int]int)
	B := make(map[int]int)

	var a, b, c int
	for i := 0; i < N; i++ {
		a = nextInt()

		_, isThere := A[a]

		if !isThere {
			A[a] = 1
		} else {
			A[a] = A[a] + 1
		}
	}

	for i := 0; i < N; i++ {
		b = nextInt()

		_, isThere := A[b]

		if !isThere {
			B[i] = 0
		} else {
			B[i] = A[b]
		}
	}

	var count int = 0

	for i := 0; i < N; i++ {
		c = nextInt()
		count += B[c-1]
	}

	fmt.Printf("%d\n", count)

}
