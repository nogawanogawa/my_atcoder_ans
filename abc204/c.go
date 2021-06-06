package main

import (
	"fmt"
)

type City struct {
	next []int
}

type Travel struct {
	dict  map[int]int
	count int
}

func dfs(i int, city []City, travel Travel) Travel {
	c_list := city[i].next

	for i := 0; i < len(c_list); i++ {
		_, ok := travel.dict[c_list[i]]
		if ok {
			continue
		} else {
			travel.dict[c_list[i]] = 1
			travel.count++
			travel = dfs(c_list[i], city, travel)
		}
	}
	return travel
}

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	city := make([]City, N)

	var A, B int
	for i := 0; i < M; i++ {
		fmt.Scanf("%d %d", &A, &B)
		city[A-1].next = append(city[A-1].next, B-1)
	}

	var count int

	for i := 0; i < N; i++ {
		travel := Travel{}
		travel.dict = map[int]int{}

		travel.dict[i] = 1
		travel.count++

		count += dfs(i, city, travel).count
	}

	fmt.Printf("%d\n", count)
}
