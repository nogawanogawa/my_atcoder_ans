package main

import (
	"fmt"
)

func main(){
	var N int
	fmt.Scanf("%d",&N)

	m := map[string]int{}
	
	var p_name string

	for i:=0; i<N; i++{
		fmt.Scanf("%s",&p_name)

		_, isThere := m[p_name]

		if !isThere {
			m[p_name] = 1
		}
	}

	fmt.Printf("%d\n",len(m))
}
