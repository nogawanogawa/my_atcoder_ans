package main

import (
    "fmt"
)

func main() {
    var S string
    fmt.Scanf("%s", &S)

    S = S[len(S)-2:] + S[:len(S)-2]

    fmt.Printf("%s\n", S)

}