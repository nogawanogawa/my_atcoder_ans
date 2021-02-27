package main

import (
	"fmt"
)

func main() {
    var N int64
    fmt.Scanf("%d", &N)

	var i, v int64

	m := make(map[int64]int)

	for i=2; i*i<=N; i++{
		v = i
		for j:=0; ; j++{
			v = v * i
			if v <= N{
				_, isThere := m[v]
				if !isThere {
					m[v] = 1
				}
			} else {
				break
			}
		}
	}

	result := N - int64(len(m))

	fmt.Printf("%d\n", result)	
}
