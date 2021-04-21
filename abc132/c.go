package main

import (
	"fmt"
	"sort"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	d := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &d[i])
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})

	l := N/2 - 1
	r := N / 2

	count := d[r] - d[l]
	fmt.Printf("%d\n", count)
}
