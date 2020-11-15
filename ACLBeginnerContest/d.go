package main

import (
    "fmt"
)

func main() {

	var n, k int
    fmt.Scanf("%d %d", &n, &m)

    A := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &A[i])
	}

	c := make([][]int, m)
	var group int = 0
	for i:=1; i<m; i++ {
		for j:=0; j<group ; j++{
			if contains(c[j], D[m][0]){
				c[j] = append(c[j], D[m][1])
				none = false
			} else if contains(c[j], D[m][1]){
				c[j] = append(c[j], D[m][0])
				none = false				
			}
		}
		if none {
			c[group] = append(c[group], D[m][0])
			c[group] = append(c[group], D[m][1])
			group += 1
		}
	}

	for g 

    fmt.Printf("%d\n", D)
}
