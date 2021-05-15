package main

import (
	"fmt"
	"sort"
)

type mountain struct {
	S string
	T int
}

func main() {

	var N int
	fmt.Scanf("%d", &N)

	m := make([]mountain, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %d", &m[i].S, &m[i].T)
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i].T > m[j].T
	})

	fmt.Printf("%s\n", m[1].S)

}
