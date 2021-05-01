package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	s := strings.Split(S, "")

	var count int
	for i := 0; i < len(S)-3; i++ {
		if (s[i] == "Z") && (s[i+1] == "O") && (s[i+2] == "N") && (s[i+3] == "e") {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
