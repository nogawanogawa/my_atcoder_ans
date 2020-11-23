package main

import (
	"fmt"
)

func main() {
    var A, B int
    fmt.Scanf("%d %d", &A, &B)
	
	var A_min, A_max float64 
	A_min = float64(A * 100) / 8
	A_max = float64((A+1) * 100) / 8
	B_min :=  B * 10 
	B_max :=  (B+1) * 10 

	var ans int

	if (A_min >= float64(B_min)) && (A_min < float64(B_max)) && (float64(B_max) - A_min >= 1){
		if A_min - float64(int(A_min)) > 0{
			ans = int(A_min) +1
		} else {
			ans = int(A_min)
		}
	} else if (A_min <= float64(B_min)) && (float64(B_min) < A_max) && (A_max - float64(B_min) >= 1){
		ans = B_min
	} else {
		ans = -1
	}

	fmt.Printf("%d\n", ans)
	
}
