package main
import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	var A int
	for i:=0; i<M; i++{
		fmt.Scanf("%d", &A)
		N = N - A
	}

	if N < 0{
		fmt.Printf("-1\n")
	} else {
		fmt.Printf("%d\n", N)
	}

}