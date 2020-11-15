package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

	point := make([][]int64, n)
	for i := 0; i < n; i++ {
		point[i] = make([]int64, 2)
		fmt.Scanf("%d %d", &point[i][0], &point[i][1])
	}

	var result bool = false
	for i:=0; i<n; i++{
		for j:=i+1; j<n; j++{
			for k:=j+1; k<n; k++{

				a_b := (point[j][1] - point[i][1]) * (point[k][0] - point[i][0])
				a_c := (point[k][1] - point[i][1]) * (point[j][0] - point[i][0])

				b_a := (point[i][1] - point[j][1]) * (point[k][0] - point[j][0])
				b_c := (point[k][1] - point[j][1]) * (point[i][0] - point[j][0])

				c_a := (point[i][1] - point[k][1]) * (point[j][0] - point[k][0])
				c_b := (point[j][1] - point[k][1]) * (point[i][0] - point[k][0])

				if (a_b == a_c) || (b_a == b_c) || ( c_a == c_b){
					result = true
				}
			}
		}
	}

	if result {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}
