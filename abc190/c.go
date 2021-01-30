package main

import (
	"fmt"
)



func check(n []int, A []int, B []int) int{
	
	var count = 0
	for i:=0; i<len(A); i++{
		if (n[A[i]-1] > 0)  && (n[B[i]-1] > 0){
			count++
		}	
	}
	return count
}

func search(i int, N int, n []int, A []int, B []int, C []int, D []int) int{

	if i >= len(C){
		return check(n, A, B)
	} else {
		m := make([]int, N)
		copy(m, n)
		m[C[i]-1] ++	
		a := search(i+1, N, m, A, B, C, D)
		
		l := make([]int, N)
		copy(l, n)
		l[D[i]-1] ++	
		b := search(i+1, N, l, A, B, C, D)

		if a > b{
			return a
		} else {
			return b
		}
	}
}

func main() {
    var N, M int
    fmt.Scanf("%d %d", &N, &M)
	
	A := make([]int, M)
	B := make([]int, M)

	for i:=0; i<M; i++{
		fmt.Scanf("%d %d", &A[i], &B[i])
	}

    var K int
    fmt.Scanf("%d", &K)

	C := make([]int, K)
	D := make([]int, K)

	for i:=0; i<K; i++{
		fmt.Scanf("%d %d", &C[i], &D[i])
	}

	n := make([]int, N)

	ans := search(0, N, n, A, B, C, D)

	fmt.Printf("%d\n", ans)
	
}
