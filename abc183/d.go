package main

import (
	"fmt"
)


func main() {

	var N int
	var W int
    fmt.Scanf("%d %d", &N, &W)

	List := make([]int, 200002)

	var S, T int
	var P int

	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d %d", &S, &T, &P)

		List[S] += P
		List[T] -= P

	}

	for i := 0; i < 200001; i++ {
		List[i+1] += List[i]
	}

	for i := 0; i < 200001; i++ {
		if List[i] > W{
			fmt.Printf("No\n" )
			return;
		}
	}

	fmt.Printf("Yes\n" )

}
