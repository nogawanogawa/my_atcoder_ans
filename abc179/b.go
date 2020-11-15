package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

	D := make([][]int, n)
	for i := 0; i < n; i++ {
		D[i] = make([]int, 2)
		fmt.Scanf("%d %d", &D[i][0], &D[i][1])
	}

	var count int = 0
	var result bool = false

	for i := 0; i < n ; i ++{
		if D[i][0] == D[i][1]{
			count = count+1 
		} else {
			count = 0
		}

		if count == 3{
			result = true
			break
		}
	}

	var msg string

	if result {
		msg = "Yes"
	} else {
		msg = "No"
	}

    fmt.Printf("%s\n", msg)
}
