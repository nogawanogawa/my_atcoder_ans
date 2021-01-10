package main

import (
	"fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

	var imax int = 1 
	for i:=0; i<N; i++{
		imax *= 2
	}

	A := make([]int64, imax)

	for i:=0; i<imax; i++{
		fmt.Scanf("%d", &A[i])		
	}

	var left int64 = 0
	var left_index int 
	for i:=0; i<imax/2; i++{
		if left < A[i]{
			left = A[i]
			left_index = i
		}
	}

	var right int64 = 0
	var right_index int 
	for i:=imax/2; i<imax; i++{
		if right < A[i]{
			right = A[i]
			right_index = i
		}
	}

	if left > right{
		fmt.Printf("%d\n", right_index+1)
	} else {
		fmt.Printf("%d\n", left_index+1)
	}
	
}
