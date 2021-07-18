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

func solve(dict map[int]int, c []int, start int, end int, old_count int) int {

	var count int
	pop := c[start-1]
	push := c[end]

	if push == pop {
		count = old_count
	} else {
		count = old_count

		_, ok := dict[push]
		if (ok) && (dict[push] > 0) {
			dict[push] += 1
		} else {
			dict[push] = 1
			count += 1
		}

		dict[pop] -= 1
		if dict[pop] == 0 {
			count -= 1
		}
	}

	return count
}

func main() {

	sc.Split(bufio.ScanWords)

	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	c := make([]int, N)
	for i := 0; i < N; i++ {
		c[i] = nextInt()
	}

	dict := make(map[int]int)

	var count int = 0
	for i := 0; i < K; i++ {
		_, ok := dict[c[i]]
		if ok {
			dict[c[i]] += 1
		} else {
			dict[c[i]] = 1
			count += 1
		}
	}

	var max_count int = count

	for i := 1; i < N-K+1; i++ {
		count = solve(dict, c, i, i+K-1, count)
		if max_count < count {
			max_count = count
		}
	}

	fmt.Printf("%d\n", max_count)

}
