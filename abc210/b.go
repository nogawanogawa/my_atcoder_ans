package main

import (
	"fmt"
	"strings"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	var S string
	fmt.Scanf("%s", &S)

	s := strings.Split(S, "")

	for i := 0; i < N; i = i + 2 {
		if s[i] == "1" {
			fmt.Printf("Takahashi\n")
			break
		}

		if s[i+1] == "1" {
			fmt.Printf("Aoki\n")
			break
		}

	}

}
