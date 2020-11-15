package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

	var count int = 0
	for a:=1; a<n ; a++{
		var b_max = int((n-1) / a)
		for b:=1; b<=b_max; b++{
			if a * b < n{
				count = count+1
			}
		}
	}

    fmt.Printf("%d\n", count)
}
