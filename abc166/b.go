package main
import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	S := make([]bool, N)

	var d, A int
	for i:=0; i<K; i++{
		fmt.Scanf("%d", &d)
		for j:=0; j<d; j++{
			fmt.Scanf("%d", &A)
			S[A-1] = true
		}
	}

	var count int = 0
	for i:=0; i<N; i++{
		if S[i] == false {
			count ++
		}
	}

	fmt.Printf("%d\n", count)

}