package main

import (
    "fmt"
)

func main() {
    var a, b, c, d uint64
    fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	
	var s string
	if (a<=d && c <= b) || (c <= b && a<=d) {
		s = "Yes"
	} else {
		s = "No"
	}

    fmt.Printf("%s\n", s)
}