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

type City struct {
	next  []int
	cost  int
	count int
}

func bfs(city []City, id int, visited []int, queue []int) {

	visited[id] = 1
	queue = queue[1:]

	next := city[id].next

	for j := 0; j < len(next); j++ {
		i := next[j]

		if visited[i] == 1 {
			continue
		} else {
			if visited[i] == 0 {
				visited[i] = 2
				queue = append(queue, i)
			}

			if city[i].cost == 0 {
				city[i].cost = city[id].cost + 1
				city[i].count = city[id].count
			} else if city[i].cost > city[id].cost+1 {
				city[i].cost = city[id].cost + 1
				city[i].count = city[id].count
			} else if city[i].cost == city[id].cost+1 {
				city[i].count = (city[i].count + city[id].count) % 1000000007
			}
		}
	}

	if len(queue) != 0 {
		bfs(city, queue[0], visited, queue)
	}
}

func main() {

	sc.Split(bufio.ScanWords)

	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	city := make([]City, N)

	var a, b int
	for i := 0; i < M; i++ {
		a, b = nextInt()-1, nextInt()-1

		city[a].next = append(city[a].next, b)
		city[b].next = append(city[b].next, a)
	}

	visited := make([]int, N)
	var queue []int
	queue = append(queue, 0)
	city[0].count = 1

	bfs(city, 0, visited, queue)

	fmt.Printf("%d\n", city[N-1].count%1000000007)
}
