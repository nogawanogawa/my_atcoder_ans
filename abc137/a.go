package main

import (
	"fmt"
)

func main() {
	var A, B int

	fmt.Scanf("%d %d", &A, &B)

	AplusB := A + B
	AminusB := A - B
	AB := A * B
	if (AplusB >= AminusB) && (AplusB >= AB){
		fmt.Printf("%d\n", AplusB)
	} else if (AminusB >= AplusB) && (AminusB >= AB){
		fmt.Printf("%d\n", AminusB)
	} else {
		fmt.Printf("%d\n", AB)
	}

}
