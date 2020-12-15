package main
import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	var count int = 0
	for i:=1; i<=N; i++{
		if !((i%3 == 0) || (i%5 == 0)){
			count += i
		}
	}

	fmt.Printf("%d\n", count)

}