package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	v := make([]float64, N)
	for i:=0; i<N; i++{
		fmt.Scanf("%f", &v[i])
	}

	sort.Slice(v, func(i, j int) bool {
        return v[i] < v[j]
    })

	var ans float64 = v[0]

	for i:=1; i<N; i++{
		ans = (ans + v[i]) / 2
	} 

	fmt.Printf("%f\n", ans)
}
