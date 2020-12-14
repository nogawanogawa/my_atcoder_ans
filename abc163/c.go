package main

import (
	"fmt"
)

func main(){
	var N int
	fmt.Scanf("%d",&N)

	m := map[int]int{}
	
	var A int 
	for i:=1; i<N; i++{
		fmt.Scanf("%d",&A)

		_, isThere := m[A]

		if !isThere {
			m[A] = 1
		} else {
			m[A] += 1
		}
	}

	for i:=1; i<=N; i++{
		_, isThere := m[i]

		if !isThere {
			fmt.Printf("0\n")
		} else {
			fmt.Printf("%d\n", m[i])
		}
	}

}
