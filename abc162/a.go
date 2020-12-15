package main

import (
	"fmt"
	"strings"
)

func main(){
	var N string
	
	fmt.Scanf("%s",&N)
	n := strings.Split(N, "")

	var flag bool = false
	for i:=0; i<len(n); i++{
		if n[i] == "7"{
			flag = true
		}
	}

	if flag {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}
