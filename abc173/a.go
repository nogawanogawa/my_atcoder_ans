package main

import (
	"fmt"
)

func main(){
	var N int
	fmt.Scanf("%d",&N)


	if  N % 1000 == 0{
		fmt.Printf("0\n")
	} else {
		n := N / 1000
		fmt.Printf("%d\n", 1000*(n+1) - N)
	}
}
