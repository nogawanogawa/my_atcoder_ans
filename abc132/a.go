package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string

	fmt.Scanf("%s", &S)
	s := strings.Split(S, "")

	m := make(map[string]int)

	var l []string

	for i := 0; i < 4; i++ {
		_, isThere := m[s[i]]

		if !isThere {
			m[s[i]] = 1
			l = append(l, s[i])
		} else {
			m[s[i]] = m[s[i]] + 1
		}
	}

	if len(l) == 2 {
		if m[l[0]] == 2 {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	} else {
		fmt.Printf("No\n")
	}

}
