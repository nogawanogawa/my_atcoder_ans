package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n, q int
    fmt.Scanf("%d %d", &n, &q)

	s := make([]rune, n)

	for i:=0; i<n ; i++{
		s[i] = '1'
	}

	var l, r int
	var d rune

    for i := 0; i<q; i++ {
        fmt.Scanf("%d %d %s", &l, &r, &d)
		for j:=l-1 ; j<r; j ++{
			s[j] = d
		}
		
		var num int64
		str := string(s)
		num, _ = strconv.ParseInt(str, 10, 64)
		result := num%998244353
	
		fmt.Printf("%d\n", result)	
	}

}
