package main

import (
    "fmt"
	"math"	
)

func main() {
	
	var R, X, Y int
	fmt.Scanf("%d %d %d",&R, &X, &Y)

	distance := math.Sqrt(float64((X * X) + (Y * Y)))

	step := int(distance / float64(R))

	if float64(step * R) == distance{
		fmt.Printf("%d\n", step)
	} else if step == 0{
		fmt.Printf("%d\n", step+2)
	} else {
		fmt.Printf("%d\n", step+1)
	}
}