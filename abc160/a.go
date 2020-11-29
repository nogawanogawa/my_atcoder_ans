package main

import (
	"fmt"
	"strings"
)

func main(){
	var S string
	fmt.Scanf("%s",&S)

	s := strings.Split(S, "")

	var ans string

	if (s[2] == s[3]) && (s[4] == s[5]){
		ans = "Yes"
	} else {
		ans = "No"
	}

	fmt.Printf("%s\n", ans)
}
