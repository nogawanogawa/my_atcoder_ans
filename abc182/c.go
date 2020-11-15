package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
    var N string
    fmt.Scanf("%s", &N)

	arr := strings.Split(N, "")

	var sum, one, two int = 0, 0, 0
	for i:=0; i<len(arr); i++{
		
		var num int
		num, _ = strconv.Atoi(arr[i])
		
		div := num % 3 
		if div == 1{
			one +=1
		}else if div == 2{
			two += 1
		}
		sum += div 
	}

	sum = sum % 3

	var ans int = -1
	
	if sum == 0{
		ans = 0
	}else if sum == 1{
		if one > 0 && len(arr)>1{
			ans = 1
		}else if two > 1 && len(arr)>2{
			ans = 2
		}
	}else if sum == 2{
		if two > 0 && len(arr)>1{
			ans = 1
		}else if one > 1 && len(arr)>2{
			ans = 2
		}
	}

	fmt.Printf("%d\n", ans)
}
