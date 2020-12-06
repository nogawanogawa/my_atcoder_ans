package main

import (
	"fmt"
)

func main(){
	var X, N int
	fmt.Scanf("%d %d",&X ,&N)

	Y := make([]bool, 102)
	var p int
	for i:=0; i<N; i++{
		fmt.Scanf("%d",&p)
		Y[p] = true
	}

	for i:=0; i<100; i++{
		if !Y[X-i] {
			fmt.Printf("%d\n", X-i)			
			break
		}
		if !Y[X+i] {
			fmt.Printf("%d\n", X+i)		
			break
		}
	}


}
