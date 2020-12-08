package main
import (
	"fmt"
)

func main() {
	var K int
	fmt.Scanf("%d", &K)

	var S string
	fmt.Scanf("%s", &S)

	if len(S) <= K {
		fmt.Printf("%s\n", S)		
	} else {
		S = S[:K] + "..."
		fmt.Printf("%s\n", S)
	}

}