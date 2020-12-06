package main
import (
	"fmt"
)

func main() {
	var X, Y int
	fmt.Scanf("%d %d", &X, &Y)

	var flag bool = false
	for i:=0; i<=X; i++{
		if Y == (i * 2 + (X - i) * 4) {
			flag = true
		}
	} 

	if flag {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}