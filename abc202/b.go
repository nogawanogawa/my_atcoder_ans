package main

import (
	"fmt"
	"strings"
)

func main() {

	var S string
	fmt.Scanf("%s", &S)

	s := strings.Split(S, "")

	var ans []string

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == "0" {
			fmt.Printf("0")
		} else if s[i] == "1" {
			ans = append(ans, "1")
			fmt.Printf("1")
		} else if s[i] == "6" {
			fmt.Printf("9")
		} else if s[i] == "8" {
			fmt.Printf("8")
		} else if s[i] == "9" {
			fmt.Printf("6")
		}
	}

	fmt.Printf("\n")

}
