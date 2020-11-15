package main

import (
    "fmt"
)

func main() {
    var s string
    fmt.Scan(&s)

    var t string
    fmt.Scan(&t)
    
	if s == "Y"{
        if t == "a"{
            t = "A"
        } else if t == "b"{
            t = "B"
        } else if t == "c"{
            t = "C"
        }
    }

    fmt.Printf("%s\n", t)
}