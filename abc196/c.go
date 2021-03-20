package main

import (
    "fmt"
	"strconv"
)

func main() {
	
	var N int64
	fmt.Scanf("%d",&N)

	var v, ans int64
	var l int
	var i_s string
	for i:=1; ; i++{

		i_s = strconv.Itoa(i)
		l = len(i_s)

		v = 1
		for j:=0; j<l; j++ {
			v = v * 10
		}


		if int64(i) + v * int64(i) > N{
			break
		} else {
			ans = int64(i)
		}
	}

	fmt.Printf("%d\n", ans)
}