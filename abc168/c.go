package main

import (
	"fmt"
	"math"
)

func main(){
	var A, B float64
	var H, M int
	
	fmt.Scanf("%f %f %d %d",&A ,&B, &H, &M)

	D := ( math.Pi * (60 * float64(H) - 11 * float64(M))) / float64(360)

	var result float64
	result = math.Sqrt(A*A + B*B - 2*A*B*math.Cos(D))

	fmt.Printf("%.10f\n",result)
}
