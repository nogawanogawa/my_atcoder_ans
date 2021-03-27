package main

import (
    "fmt"
	"strings"
)

func main() {
	
	var H, W, X, Y int
	fmt.Scanf("%d %d %d %d",&H, &W, &X, &Y)

	S := make([][]string, H)
	for i:=0; i<H; i++{
		S[i] = make([]string, W)
	}

	var s string
	for i:=0; i<H; i++{
		fmt.Scanf("%s",&s)
		s_ := strings.Split(s, "")
		for j:=0; j<W; j++{
			S[i][j] = s_[j]
		}
	}

	var count int = 1
	for i:=Y-2; i>=0; i--{
		if S[X-1][i] == "."{
			count += 1
		} else {
			break
		}
	}

	for i:=Y; i<W; i++{
		if S[X-1][i] == "."{
			count += 1
		} else {
			break
		}
	}

	for i:=X-2; i>=0; i--{
		if S[i][Y-1] == "."{
			count += 1
		} else {
			break
		}
	}

	for i:=X; i<H; i++{
		if S[i][Y-1] == "."{
			count += 1
		} else {
			break
		}
	}

	fmt.Printf("%d\n", count)
}
