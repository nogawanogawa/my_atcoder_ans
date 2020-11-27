package main

import (
	"fmt"
)

func main(){
	var S string
	fmt.Scanf("%s",&S)

	var ans int
	if S == "RRR" {
		ans = 3
	} else if (S == "RRS") || (S == "SRR") {
		ans = 2
	} else if S == "SSS"{
		ans = 0
	} else {
		ans = 1
	}

	fmt.Printf("%d\n", ans)
}
