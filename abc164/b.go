package main
import (
	"fmt"
)

func main() {
	var A, B, C, D int
	fmt.Scanf("%d %d %d %d", &A, &B, &C, &D)

	var flag bool
	for i:=0; ; i++{
		C = C-B
		if C <= 0 {
			flag = true
			break
		}

		A = A-D
		if A <= 0 {
			flag = false
			break
		}
	}

	if flag{
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}

}