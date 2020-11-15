package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)
	
	var a, b int64
	var result int64
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a, &b)	
		result += (b+a)*(b-a+1)/2
	}

	fmt.Printf("%d\n", result)
	
}
