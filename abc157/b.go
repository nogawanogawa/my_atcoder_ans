package main

import (
    "fmt"
)

func main() {
	var A [3][3]int
    fmt.Scanf("%d %d %d", &A[0][0], &A[0][1], &A[0][2])
    fmt.Scanf("%d %d %d", &A[1][0], &A[1][1], &A[1][2])
    fmt.Scanf("%d %d %d", &A[2][0], &A[2][1], &A[2][2])

	var bingo [3][3]bool

	var N int
	fmt.Scanf("%d", &N)

	var b int
	for i:=0; i<N; i++{
		fmt.Scanf("%d", &b)
		for j:=0; j<3; j++{
			for k:=0; k<3; k++{
				if A[j][k] == b {
					bingo[j][k] = true
				}
			}			
		}
	}

	ans := "No" 
	for i:=0; i<3; i++{
		if bingo[i][0] && bingo[i][1] && bingo[i][2]{
			ans = "Yes"
		} 
		if bingo[0][i]&& bingo[1][i] && bingo[2][i]{
			ans = "Yes"
		} 
	}

	if bingo[0][0] && bingo[1][1] && bingo[2][2]{
		ans = "Yes"
	} 

	if bingo[2][0] && bingo[1][1] && bingo[0][2]{
		ans = "Yes"
	} 

	fmt.Printf("%s\n", ans)	
}