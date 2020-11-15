package main

import (
	"fmt"
)


func main() {
    var N int
    fmt.Scanf("%d", &N)

	var ans int =0
	var max_right int = 0
	var p_now int = 0
	var total int = 0

	var num int

	for i:=0; i<N; i++{
		fmt.Scanf("%d", &num)

		total = total + num 

		if max_right < total { 
			max_right = total
		}

		if ans < p_now + max_right {
			ans = p_now + max_right
		}
		
		p_now = p_now + total
	}

	fmt.Printf("%d\n", ans)
}
