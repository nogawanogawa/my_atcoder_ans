package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

	var count int = 0
	for i:=1; i<=N; i++{
		if ((0 <= i) && (i <= 9)) || ((100 <= i) && (i <= 999)) || ((10000 <= i) && (i <= 99999)){
			count ++ 
		}
	}    
	fmt.Printf("%d\n", count)
}
