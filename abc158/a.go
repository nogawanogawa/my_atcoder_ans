package main

import (
    "fmt"
)

func main() {
    var S string 
    fmt.Scanf("%s", &S)

    var ans string
    if (S == "AAA") || (S == "BBB"){
        ans = "No"
    } else {
        ans = "Yes"
    }

	fmt.Printf("%s\n", ans)
	
}