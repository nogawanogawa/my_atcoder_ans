package main

import (
    "fmt"
)

func main() {
    var k int
    fmt.Scanf("%d", &k)
	
	var s string = ""
	for i := 0; i<k; i++{
		s += "ACL"
	}

    fmt.Printf("%s\n", s)
}