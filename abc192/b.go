package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsFirstUpper(v string) bool {
	if v == "" {
	   return false;
	}
	r := rune(v[0])
	return unicode.IsUpper(r)
} 

func main() {
	var S string

	fmt.Scanf("%s", &S)
	s := strings.Split(S, "")

	var ans bool = true
	for i:=0; i<len(s); i++{
		check := IsFirstUpper(s[i])
		if i % 2 == 0{
			if check{
				ans = false
			}
		} else {
			if !check{
				ans = false
			}
		}
	}

	if ans {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}
