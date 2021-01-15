package main

import (
    "fmt"
)

func main() {

    var S string
    fmt.Scanf("%s", &S)

    if S == "SUN"{
        fmt.Printf("7\n")
    } else if S == "MON"{
        fmt.Printf("6\n")
    } else if S == "TUE"{
        fmt.Printf("5\n")
    } else if S == "WED"{
        fmt.Printf("4\n")
    } else if S == "THU"{
        fmt.Printf("3\n")
    } else if S == "FRI"{
        fmt.Printf("2\n")
    } else if S == "SAT"{
        fmt.Printf("1\n")
    }
}