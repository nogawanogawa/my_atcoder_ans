package main

import (
	"fmt"
	"strings"
)

func find_tail_r(c []string, point int) int{
	//末尾のRを見つける
	var result int
	for i:=point; i>0; i--{
		if c[i] == "R"{
			result = i
			break
		}
	}

	return result
}

func main(){
	var N int
	fmt.Scanf("%d",&N)

	var C string
	fmt.Scanf("%s",&C)
	c := strings.Split(C, "")

	var point int
	point = find_tail_r(c, N-1)
	
	var count int = 0
	for i:=0; i<point; i++{
		if c[i] == "W"{
			count += 1
			c[point] = "W"
			c[i] = "R"

			point = find_tail_r(c, point-1)
		}
	}

	fmt.Printf("%d\n", count)
}
