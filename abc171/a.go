package main

import (
	"fmt"
	"strings"
)

func main(){
	var alpha string
	fmt.Scanf("%s",&alpha)

	if alpha == strings.ToLower(alpha) {
		fmt.Printf("a\n")
	} else {
		fmt.Printf("A\n")
	}

}
