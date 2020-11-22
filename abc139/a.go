package main

import (
	"fmt"
	"strings"
)

func main() {
	var S, T string

	fmt.Scanf("%s", &S)

	s := strings.Split(S, "")

	fmt.Scanf("%s", &T)

	t := strings.Split(T, "")

	var count int = 0
	for i:=0; i<3; i++{
		if s[i] == t[i]{
			count += 1
		}
	}

    fmt.Printf("%d\n", count)
}
