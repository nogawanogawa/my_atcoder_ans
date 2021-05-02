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

type Node struct {
	next int
	pass int
}

func dfs(nodes []Node, j int) int {

	next := nodes[j].next

	if nodes[next].pass == 1 {
		nodes[j].pass = 2
		return 1
	} else if nodes[next].pass == 2 {
		nodes[j].pass = 2
		return 0
	} else {
		nodes[j].pass = 1
		count := dfs(nodes, next)

		nodes[j].pass = 2
		return count
	}

}

func powMod(a int, n int, b int) int {
	if n == 1 {
		return a
	}
	if n%2 == 1 {
		return a * powMod(a, n-1, b) % b
	}
	ret := powMod(a, n/2, b)
	ret = (ret * ret) % b

	return ret
}

func main() {

	sc.Split(bufio.ScanWords)

	var N int
	fmt.Scanf("%d", &N)

	nodes := make([]Node, N)

	for i := 0; i < N; i++ {
		nodes[i].next = nextInt() - 1
	}

	var count int

	for i := 0; i < N; i++ {
		if nodes[i].pass == 0 {
			count += dfs(nodes, i)
		}
	}

	var ans int
	ans = powMod(2, count, 998244353)

	if ans == 0 {
		ans = 998244352
	} else {
		ans = ans - 1
	}

	fmt.Printf("%d\n", ans)
}
