package main

import (
	"fmt"
)

func main() {
	var N float64
	fmt.Scanf("%f", &N)

	ans := int(N * 1.08)
	if ans == 206 {
		fmt.Printf("so-so\n")
	} else if ans < 206 {
		fmt.Printf("Yay!\n")
	} else {
		fmt.Printf(":(\n")
	}

}
