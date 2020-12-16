package main
import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	A := make([]int, N)
	var sum, count int 

	for i:=0; i<N; i++{
		fmt.Scanf("%d", &A[i])
		sum += A[i]
	}

	var threshold float64
	threshold = float64(sum) / (4 * float64(M))

	for i:=0; i<N; i++{
		if float64(A[i]) >= threshold{
			count ++
		}
	}

	if count >= M{
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}

}