package main

import (
	"fmt"
)

func main() {
	var P, Q, R int

	fmt.Scanf("%d %d %d", &P, &Q, &R)

	if (P >= Q) && (P >= R) {
		fmt.Printf("%d\n", Q+R)
	} else if (Q >= P) && (Q >= R) {
		fmt.Printf("%d\n", P+R)
	} else {
		fmt.Printf("%d\n", Q+P)
	}
}
