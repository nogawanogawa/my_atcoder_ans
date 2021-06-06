package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)

	if x == 0 {
		if y == 0 {
			fmt.Printf("%d\n", 0)
		} else if y == 1 {
			fmt.Printf("%d\n", 2)
		} else if y == 2 {
			fmt.Printf("%d\n", 1)
		}
	} else if x == 1 {
		if y == 0 {
			fmt.Printf("%d\n", 2)
		} else if y == 1 {
			fmt.Printf("%d\n", 1)
		} else if y == 2 {
			fmt.Printf("%d\n", 0)
		}
	} else if x == 2 {
		if y == 0 {
			fmt.Printf("%d\n", 1)
		} else if y == 1 {
			fmt.Printf("%d\n", 0)
		} else if y == 2 {
			fmt.Printf("%d\n", 2)
		}
	}
}
