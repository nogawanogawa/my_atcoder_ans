package main
import (
	"fmt"
)

func main() {
	var X,i int64
	fmt.Scanf("%d", &X)

	var result int = 0 
	for i=100; i<X ; i+=i/100{
		result ++
	}

	fmt.Printf("%d\n", result)
}