package main

import (
	"fmt"
	"strings"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

	S_a := make(map[string]int, N)
	S_b := make(map[string]int, N)

	var tmp string
	var ans string = ""

	for i:=0; i<N; i++{
		fmt.Scanf("%s", &tmp)
		t := strings.Split(tmp, "")
		if t[0] == "!"{
			S_b[tmp[1:]] = 1

			if _, ok := S_a[tmp[1:]]; ok {
				ans = tmp[1:]
			}

		} else {
			S_a[tmp] = 1

			if _, ok := S_b[tmp]; ok {
				ans = tmp
			}

		}
	}

	if ans != ""{
		fmt.Printf("%s\n", ans)
	} else {
		fmt.Printf("satisfiable\n")
	}
	
}
