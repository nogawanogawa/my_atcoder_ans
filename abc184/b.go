package main

import (
	"fmt"
	"strings"
)

func main() {
    var N, X int
    fmt.Scanf("%d %d", &N, &X)
	
	var S string
    fmt.Scanf("%s", &S)

	result := strings.Split(S, "")

	for i:=0; i<N; i++{
		if result[i] == "o"{
			X += 1
		} else if X > 0{
			X -= 1
		}
	}

	fmt.Printf("%d\n", X)
	
}
