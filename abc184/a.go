package main

import (
    "fmt"
)

func main() {
    var a, b, c, d int 
    fmt.Scanf("%d %d", &a, &b)
    fmt.Scanf("%d %d", &c, &d)

    ans := a * d - b * c

	fmt.Printf("%d\n", ans)
	
}