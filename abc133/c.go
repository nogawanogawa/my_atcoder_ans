package main

import (
	"fmt"
)

func main() {

	var L, R int
	fmt.Scanf("%d %d", &L, &R)

	var min int = 2020
	if R-L > 2019 {
		fmt.Printf("%d\n", 0)
	} else {
		for l := L; l <= R-1; l++ {
			for r := l + 1; r <= R; r++ {
				if (l*r)%2019 < min {
					min = (l * r) % 2019
				}
			}
		}
		fmt.Printf("%d\n", min)
	}

}
