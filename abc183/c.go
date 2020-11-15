package main

import (
	"fmt"
)

func remove(t []int, i int) []int {
	s := make([]int, len(t))

	copy(s, t)
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func greedy(n int, k int64, list []int, T [][]int64, now int) int{
	

	var num int = 0
	if n == 1{
		if T[now-1][0] == k{
			return 1
		} else {
			return 0
		}
	} else if  k<0{
		return 0
	}

	for i:=1; i<n; i++{
		next := list[i]
		var list_2 = remove(list, i)

		w := T[now-1][next-1]

		num += greedy(n-1, k-w, list_2, T, next)
	}

	return num
}

func main() {
	var N int
	var K int64
    fmt.Scanf("%d %d", &N, &K)

	T := make([][]int64, N)
	for i := 0; i < N; i++ {
		T[i] = make([]int64, N)
		for j:=0; j<N; j++{
			fmt.Scanf("%d", &T[i][j])
		}
	}

	List := make([]int, N)

	for i := 0; i < N; i++ {
		List[i] = i+1
	}

	ans := greedy(N, K, List, T, 1)

	fmt.Printf("%d\n", ans)
}
