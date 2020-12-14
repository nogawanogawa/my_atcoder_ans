package main

import (
	"fmt"
	"math"
)

func main(){
	var R int
	
	fmt.Scanf("%d",&R)

	D := ( math.Pi * 2 * float64(R))

	fmt.Printf("%.10f\n", D)
}
