package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string

	fmt.Scanf("%s", &S)
	s := strings.Split(S, "")

	var tmp string = ""
	var flag bool = false
	for i := 0; i < 4; i++ {
		if tmp == s[i] {
			flag = true
			break
		} else {
			tmp = s[i]
		}
	}

	if flag {
		fmt.Printf("Bad\n")
	} else {
		fmt.Printf("Good\n")
	}

}
