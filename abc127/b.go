package main

import (
	"fmt"
)

func main() {
	var r, D, x int
	fmt.Scanf("%d %d %d", &r, &D, &x)

	for i := 2001; i <= 2010; i++ {
		x = r*x - D
		fmt.Printf("%d\n", x)
	}

}
