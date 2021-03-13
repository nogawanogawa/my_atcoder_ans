package main

import (
    "fmt"
)

func main() {
	
	var N int64
	fmt.Scanf("%d",&N)

	c_1 := N - 1000 + 1
	c_2 := N - 1000000 + 1
	c_3 := N - 1000000000 + 1
	c_4 := N - 1000000000000 + 1
	c_5 := N - 1000000000000000 + 1

	var ans int64 = 0
	if c_1 > 0{
		ans += c_1
	}
	if c_2 > 0{
		ans += c_2	
	}
	if c_3 > 0{
		ans += c_3	
	}
	if c_4 > 0{
		ans += c_4
	}
	if c_5 > 0{
		ans += c_5	
	}

	fmt.Printf("%d\n", ans)
}