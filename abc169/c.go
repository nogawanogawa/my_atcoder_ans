package main

import (
	"fmt"
)

func main(){
	var A int64
	var B float64
	fmt.Scanf("%d %f",&A ,&B)

	var result int64
	result = int64((A * int64(B * 100.0 + 0.5)) / 100)

	fmt.Printf("%d\n",result)
}
