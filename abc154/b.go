package main

import (
    "fmt"
    "strings"
)

func main() {
    var S string
    fmt.Scanf("%s", &S)

    var ans string = ""
    end := len(strings.Split(S, ""))
    for i:=0; i<end;i++{
        ans += "x"
    }

    fmt.Printf("%s\n", ans)

}