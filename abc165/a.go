package main

import (
	"fmt"
)

func main(){
	var K int
	fmt.Scanf("%d", &K)

	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	var flag bool
	for i:=A; i<=B; i++{
		if i%K == 0{
			flag = true
			break
		}
	}

	if flag {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("NG\n")
	}
}
