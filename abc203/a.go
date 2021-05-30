package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a == b {
		fmt.Printf("%d\n", c)
	} else if b == c {
		fmt.Printf("%d\n", a)
	} else if c == a {
		fmt.Printf("%d\n", b)
	} else {
		fmt.Printf("%d\n", 0)
	}

}
