package main

import (
	"fmt"
)

func main() {

	var S1, S2, S3, S4 string
	m := make(map[string]int)

	fmt.Scanf("%s", &S1)
	m[S1] = 1
	fmt.Scanf("%s", &S2)
	m[S2] = 1
	fmt.Scanf("%s", &S3)
	m[S3] = 1
	fmt.Scanf("%s", &S4)
	m[S4] = 1

	_, isThere := m["H"]
	if !isThere {
		fmt.Printf("No\n")
		return
	}

	_, isThere = m["2B"]
	if !isThere {
		fmt.Printf("No\n")
		return
	}

	_, isThere = m["3B"]
	if !isThere {
		fmt.Printf("No\n")
		return
	}

	_, isThere = m["HR"]
	if !isThere {
		fmt.Printf("No\n")
		return
	}

	fmt.Printf("Yes\n")

}
