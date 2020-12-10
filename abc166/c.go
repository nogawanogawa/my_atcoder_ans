package main

import (
	"fmt"
)

type Peak struct {
	H int 
	connect []int
}

func main(){
	var N, M, A, B int
	fmt.Scanf("%d %d",&N ,&M)

	P := make([]Peak, N)

	for i:=0; i<N; i++{
		fmt.Scanf("%d",&P[i].H)
	}

	for i:=0; i<M; i++{
		fmt.Scanf("%d %d",&A, &B)
		P[A-1].connect = append(P[A-1].connect, B-1)
		P[B-1].connect = append(P[B-1].connect, A-1)
	}

	var c, a, count int 
	var flag bool
	for i:=0; i<N; i++{
		flag = true

		c = len(P[i].connect)
		for j:=0; j<c; j++{
			a = P[i].connect[j]
			if P[i].H <= P[a].H {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}	

	fmt.Printf("%d\n",count)
}
