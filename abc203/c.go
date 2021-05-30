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

type Friend struct {
	A int64
	B int64
}

func main() {

	sc.Split(bufio.ScanWords)

	var N, K int64
	fmt.Scanf("%d %d", &N, &K)

	friend := make([]Friend, N)

	var i int64
	for i = 0; i < N; i++ {
		friend[i].A = nextInt()
		friend[i].B = nextInt()
	}

	sort.Slice(friend, func(i, j int) bool {
		return friend[i].A < friend[j].A
	})

	var money int64 = K
	for i = 0; i < N; i++ {
		if money >= friend[i].A {
			money += friend[i].B
		} else {
			break
		}
	}

	fmt.Printf("%d\n", money)

}
