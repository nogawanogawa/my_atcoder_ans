package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	next  []int
	color int
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func dfs(nodes []Node, before []int, count int, start int) {

	now := nodes[start]
	if count%2 == 0 {
		nodes[start].color = 0
	} else {
		nodes[start].color = 1
	}

	for i := 0; i < len(now.next); i++ {
		if contains(before, now.next[i]) == false {
			var before_ []int
			copy(before_, before)
			before_ = append(before_, start)
			dfs(nodes, before_, count+1, now.next[i])
		}
	}
}

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

	var N, Q int
	fmt.Scanf("%d %d", &N, &Q)

	nodes := make([]Node, N)

	var a, b int
	for i := 0; i < N-1; i++ {
		a = nextInt() - 1
		b = nextInt() - 1
		nodes[a].next = append(nodes[a].next, b)
		nodes[b].next = append(nodes[b].next, a)
	}

	C := make([]int, Q)
	D := make([]int, Q)
	for i := 0; i < Q; i++ {
		C[i] = nextInt() - 1
		D[i] = nextInt() - 1
	}

	var before []int
	dfs(nodes, before, 0, 0)

	for i := 0; i < Q; i++ {
		x := nodes[C[i]].color
		y := nodes[D[i]].color
		if x == y {
			fmt.Printf("Town\n")
		} else {
			fmt.Printf("Road\n")
		}
	}
}
