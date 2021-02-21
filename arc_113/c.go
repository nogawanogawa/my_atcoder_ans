package main

import (
	"fmt"
	"strings"
)

func main() {
    var S string
    fmt.Scanf("%s", &S)

	s := strings.Split(S, "")
	var k int =len(s)
	var l int = 0
	var count int = 0
	for i:=len(s)-3; i>=0; i--{

		if (s[i] == s[i+1]) && (s[i+1] != s[i+2]){
			for j:=i+2; j<k; j++{
				if s[j] != s[i]{
					count += 1
					s[j] = s[i]
				}
			}

			if k == len(s){
				count += 0
			} else if s[i] != s[k]{
				count += l
			}

			k = i
			l = len(s) - k
		}
	}

	fmt.Printf("%d\n", count)
}

