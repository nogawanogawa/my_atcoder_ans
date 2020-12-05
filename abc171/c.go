package main

import (
	"fmt"
)

func int_to_char(n int64) string{
	var c string 
	
	if n == 0{
		c = "z"
	} else if n == 1{
		c = "a"
	} else if n == 2{
		c = "b"
	} else if n == 3{
		c = "c"
	} else if n == 4{
		c = "d"
	} else if n == 5{
		c = "e"
	} else if n == 6{
		c = "f"										
	} else if n == 7{
		c = "g"
	} else if n == 8{
		c = "h"
	} else if n == 9{
		c = "i"
	} else if n == 10{
		c = "j"
	} else if n == 11{
		c = "k"
	} else if n == 12{
		c = "l"
	} else if n == 13{
		c = "m"
	} else if n == 14{
		c = "n"
	} else if n == 15{
		c = "o"
	} else if n == 16{
		c = "p"
	} else if n == 17{
		c = "q"
	} else if n == 18{
		c = "r"										
	} else if n == 19{
		c = "s"
	} else if n == 20{
		c = "t"
	} else if n == 21{
		c = "u"
	} else if n == 22{
		c = "v"
	} else if n == 23{
		c = "w"
	} else if n == 24{
		c = "x"
	} else if n == 25{
		c = "y"		
	}

	return c
}

func main(){
	var N int64
	fmt.Scanf("%d",&N)

	// 文字列の長さを調査
	var div int64 = 1
	var max int64 = 0
	var i, l int64 

	for i=1; i<16; i++{
		div = div * 26
		max = max + div

		if N <= max {
			l = i
			break
		}
	}

	div = 1
	var s int64
	var ans string = ""
	for i=0; i<l; i++{
		s = N % 26
		N = (N-1) / 26
		c := int_to_char(s)
		ans = c + ans
	}
	fmt.Printf("%s", ans)						

}
