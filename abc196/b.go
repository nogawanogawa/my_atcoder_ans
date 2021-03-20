package main

import (
    "fmt"
	"strings"
)

func main() {
	
	var X string
	fmt.Scanf("%s",&X)

	arr := strings.Split(X, ".")

	x := arr[0]

	fmt.Printf("%s\n", x)
}
