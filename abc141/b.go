package main

import (
    "fmt"
    "strings"
)

func main() {
    var S string
    fmt.Scanf("%s", &S)

    s := strings.Split(S, "")

    var flag bool = true
    for i:=0; i<len(s); i++{
        if i%2 == 0{
            if s[i] == "L"{
                flag = false
            }
        } else {
            if s[i] == "R"{
                flag = false
            }
        }
    }

    if flag {
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }

    
}