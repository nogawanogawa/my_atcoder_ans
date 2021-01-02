package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    var A, B string
    fmt.Scanf("%s %s", &A, &B)
    A_ := strings.Split(A, "")
    B_ := strings.Split(B, "")
    
    a_1, _ := strconv.Atoi(A_[0])
    a_2, _ := strconv.Atoi(A_[1])
    a_3, _ := strconv.Atoi(A_[2])
    a := a_1 + a_2 + a_3

    b_1, _ := strconv.Atoi(B_[0])
    b_2, _ := strconv.Atoi(B_[1])
    b_3, _ := strconv.Atoi(B_[2])
    b := b_1 + b_2 + b_3


    if a > b{
        fmt.Printf("%d\n", a)
    } else {
    	fmt.Printf("%d\n", b)
    }	
}