package main

import (
	"fmt"
)


func main() {

    var N int
    fmt.Scanf("%d", &N)

	var ans int = 0

	if N == 1{ 
		ans = 0
	} else if N == 2{
		ans = 1
	} else {
		for i:=2 ; i<N; i ++ {
			ans += i
		} 
		ans += 1
	}
	
	fmt.Printf("%d\n", ans)
}
