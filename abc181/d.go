package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
    var s string
    fmt.Scanf("%s", &s)

	nums := strings.Split(s, "")

	for i:=1; i<10; i++{
		
	} 

	var result bool = false
	if len(nums) == 1{
		num, _ := strconv.Atoi(nums[0])
		if num == 8{
			result = true
		}
	}else if len(nums) == 2{
		num_0, _ := strconv.Atoi(nums[0] + nums[1])
		num_1, _ := strconv.Atoi(nums[1] + nums[0])

		if (num_0 % 8 == 0) || (num_1 % 8 == 0){
			result = true
		}
	}else{
		for i:=112; i<1000; i+=8{
			strconv.Itoa(i)
			
		}
	}

	if result {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}
