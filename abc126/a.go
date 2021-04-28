package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	var S string
	fmt.Scanf("%s", &S)

	s := strings.Split(S, "")
	s[K-1] = strings.ToLower(s[K-1])
	S = strings.Join(s, "")

	fmt.Printf("%s\n", S)

}
