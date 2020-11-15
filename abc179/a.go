package main

import (
    "fmt"
)

func main() {
    var s string
    fmt.Scan(&s)
	
	if s[len(s)-1:] == "s"{
        s = s + "es"
    } else {
        s = s + "s"
    }

    fmt.Printf("%s\n", s)
}