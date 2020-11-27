package main

import (
	"fmt"
)

func main(){
	var N int
	fmt.Scanf("%d",&N)

	L := make([]int, N)

	for i:=0; i<N; i++{
		fmt.Scanf("%d",&L[i])
	}

	var count int = 0
	for i:=0; i<N-2; i++{
		for j:=i+1; j<N-1; j++{
			for k:=j+1; k<N; k++{
				if (L[i] != L[j]) && (L[j] != L[k]) && (L[k] != L[i]) && 
				(L[i] + L[j] > L[k]) && (L[j] + L[k] > L[i]) && (L[k] + L[i] > L[j]) {
					count +=1
				}
			}
		}
	}

	fmt.Printf("%d\n", count)
}
