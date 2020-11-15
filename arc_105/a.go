package main

import (
    "fmt"
)

func main() {
    var a,b,c,d int
    fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

    flag := false 

    flag = (a == b+c+d || a+b == c+d || a+c == b+d || a+d == b+c || a+b+c == d || a+b+d == c || a+c+d == b )

    var s string
    if flag == true{
        s = "Yes"
    }else{
        s = "No"
    }
    fmt.Printf("%s\n", s)
}