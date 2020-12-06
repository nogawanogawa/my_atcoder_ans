package main

import (
	"fmt"
)

func main(){
	var x int

	for i:=0; i<5; i++{
		fmt.Scanf("%d", &x)
		if x == 0{
			fmt.Printf("%d\n", i+1)
		}
	}

}
