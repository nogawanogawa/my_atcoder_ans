package main

import (
	"fmt"
)

func main() {
    var N int64
    fmt.Scanf("%d", &N)
	
	A := make([]int64, N)
	var i int64
	for i=0; i<N; i++{
		fmt.Scanf("%d", &A[i])
	}

	var mikan int64 = 0

	var max, min int64
	var l,r int64
	for l=0; l<N; l++{
		min = A[l]
		for r=l; r<N; r++{
			if A[r] < min {
				min = A[r]
			}
			max = min * (r - l + 1)
			if max > mikan {
				mikan = max
			}
		}
	}

	fmt.Printf("%d\n", mikan)
	
}
