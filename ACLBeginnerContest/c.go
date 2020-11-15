package main

import (
    "fmt"
)

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func main() {

	var n, m int
    fmt.Scanf("%d %d", &n, &m)

    D := make([][]int, m)
    for i := 0; i < m; i++ {
        D[i] = make([]int, 2)
        fmt.Scanf("%d %d", &D[i][0], &D[i][1])
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
