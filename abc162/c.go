package main

import (
	"fmt"
)


func gcd(a, b int) int {
	if b == 0 {
	  return a
	}
	return gcd(b, a % b)
}

func main(){
	var K int
	fmt.Scanf("%d",&K)

	var d, count int
	for a:=1; a<=K; a++{
		for b:=1; b<=K; b++{
			d = gcd(a, b)
			for c:=1; c<=K; c++{
				count += gcd(c, d)
			}
		}
	}  

	fmt.Printf("%d\n", count)

}
