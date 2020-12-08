package main

import (
	"fmt"
)

func main(){
	var N int
	fmt.Scanf("%d", &N)

	one := N % 10 

	if (one == 3) {
		fmt.Printf("bon\n")
	} else if (one == 0) || (one == 1) || (one == 6) || (one == 8){
		fmt.Printf("pon\n")
	} else {
		fmt.Printf("hon\n")
	}

}
