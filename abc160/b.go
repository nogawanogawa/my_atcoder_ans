package main

import (
	"fmt"
)

func main(){
	var X int
	fmt.Scanf("%d",&X)

	// X = 500 * a + a_
	// a_ = 5 * b + b_
	var a, b, a_ int 
	a = X / 500 
	a_ = X % 500
	b = a_ / 5
	
	ans := a * 1000 + b * 5 

	fmt.Printf("%d\n", ans)
}
