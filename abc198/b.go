package main

import (
    "fmt"
	"strconv"
	"strings"
)

func main() {
	
	var N int
	fmt.Scanf("%d",&N)

	n := strconv.Itoa(N)
	n = strings.TrimRight(n, "0")

	var flag bool = true
	for i:=0; i<len(n); i++{
		if n[i] != n[len(n)-1-i]{
			flag = false
		}
	}

	if flag {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}

}
