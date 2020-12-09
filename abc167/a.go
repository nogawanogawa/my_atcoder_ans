package main

import (
	"fmt"
)

func main(){
	var S string
	fmt.Scanf("%s", &S)

	var T string
	fmt.Scanf("%s", &T)

	if S == T[:len(T)-1]{
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}


}
