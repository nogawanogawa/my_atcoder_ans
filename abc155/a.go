package main

import (
    "fmt"
)

func main() {
    var A, B, C int
    fmt.Scanf("%d %d %d", &A, &B, &C)

    var ans string

	if ((A == B) && (C != B)) || ((B == C) && (A != B)) || ((A == C) && (C != B)){
        ans = "Yes"
    } else {
        ans = "No"
    }

	fmt.Printf("%s\n", ans)	
}