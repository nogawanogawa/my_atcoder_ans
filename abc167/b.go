package main
import (
	"fmt"
)

func main() {
	var A, B, C, K int
	fmt.Scanf("%d %d %d %d", &A, &B, &C, &K)
	var result int

	if A >= K {
		result = K
	} else  if (A+B) >= K{
		result = A
	} else if (A+B+C) >= K{
		result = A - (K - A - B) 
	}

	fmt.Printf("%d\n", result)

}