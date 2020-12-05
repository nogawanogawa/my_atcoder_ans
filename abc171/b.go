package main
import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	p := make([]int, N)
	for i:=0; i<N; i++{
		fmt.Scanf("%d", &p[i])
	}

	sort.Slice(p, func(i, j int) bool {
        return p[i] < p[j]
	})
	
	var sum int = 0
	for i:=0; i<K; i++{
		sum += p[i]
	}

	fmt.Printf("%d\n", sum)
}