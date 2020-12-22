package main

import (
	"fmt"
	"strconv"
	"strings"
)

func to_oct(n int) string {
	var oct string
	for i:=0; ;i++{
		a := n % 8
		oct = strconv.Itoa(a) + oct

		n /= 8
		if n == 0{
			break
		}
	}

	return oct
}

func main() {
    var N int
    fmt.Scanf("%d", &N)
	var decimal, oct string
	var count int = 0
	for i:=1; i<=N; i++{
		decimal = strconv.Itoa(i)
		oct = to_oct(i)
		if (!strings.Contains(decimal, "7")) && (!strings.Contains(oct, "7")){
			count++
		} 
	}

	fmt.Printf("%d\n", count)
	
}
