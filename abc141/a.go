package main

import (
    "fmt"
)

func main() {

    var S string
    fmt.Scanf("%s", &S)

    if S == "Sunny"{
        fmt.Printf("Cloudy\n")
    } else if S == "Cloudy"{
        fmt.Printf("Rainy\n")
    } else {
        fmt.Printf("Sunny\n")
    }

}