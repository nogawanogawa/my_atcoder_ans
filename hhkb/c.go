package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

	p := [200001] int{}

	min := 0
	var s int
	for i := 0; i<n; i++{
		fmt.Scanf("%d", &s)
		p[s] = 1

		for j:=min; j<2000001; j++{
			if p[min] == 1{
				min ++
			}else{
				fmt.Printf("%d\n", min)
				break
			}
		}
	}
}
