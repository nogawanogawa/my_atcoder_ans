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
	i, e := strconv.ParseInt(sc.Text(), 10, 32)
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

	dict := make(map[int]int)
	var l []int
	for i := 0; i < N; i++ {
		_, isThere := dict[A[i]%200]

		if !isThere {
			dict[A[i]%200] = 1
			l = append(l, A[i]%200)
		} else {
			dict[A[i]%200] = dict[A[i]%200] + 1
		}
	}

	var count int = 0

	for i := 0; i < len(l); i++ {
		a := dict[l[i]]
		if a > 1 {
			count += (a * (a - 1)) / 2
		}
	}

	fmt.Printf("%d\n", count)

}
