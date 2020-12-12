package main

import (
	"fmt"
)

func main(){
	var S, W int
	fmt.Scanf("%d %d", &S, &W)

	if W >= S {
		fmt.Printf("unsafe\n")
	} else {
		fmt.Printf("safe\n")
	}
}
