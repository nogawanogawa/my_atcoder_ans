package main

import (
    "fmt"
)

func main() {
    var N, D int
    fmt.Scanf("%d %d", &N, &D)

	width := 1 + 2 * D
	if N % width == 0 {
		fmt.Printf("%d\n",N/width)
	} else {
		fmt.Printf("%d\n",N/width + 1)
	}
}
