package main

import (
    "fmt"
)

func main() {
    var K, X int
    fmt.Scanf("%d %d", &K, &X)

	min := X - K + 1
	max := X + K - 1

	for i:=min; i<=max; i++{
		if (-1000000 <= i) && ( i <= 1000000){
			fmt.Printf("%d ", i)
		}
	}    
}
