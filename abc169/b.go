package main
import (
	"fmt"
)

func main() {
	var i, N, A int64
	fmt.Scanf("%d", &N)

	var result int64 = 1
	for i=0; i<N; i++{
		fmt.Scanf("%d", &A)
		
		if result > 0 {
			if (1000000000000000000 / result < A) && (A!=0){
				result = -1
			} else {
				result *= A
			}
		} else if result == 0 {
			result = 0
		} else if result < 0{
			if A == 0{
				result = 0
			} else {
				result = -1
			}
		}

		if result > 1000000000000000000{
			result = -1
		} 

	}

	if (result > 1000000000000000000) || (result < 0){
		result = -1
	}

	fmt.Printf("%d\n", result)
}