package main

import (
	"fmt"
	"strings"
)


func main() {
    var H, W int
    fmt.Scanf("%d %d", &H, &W)
	
	S := make([][]string, H)
	for i:=0; i<H; i++{
		S[i] = make([]string, W)
	}

	var row string
	for i:=0; i<H; i++{
		fmt.Scanf("%s", &row)
		char := strings.Split(row, "")
		for j:=0; j<W; j++{
			S[i][j] = char[j]
		}
	}

	var edge int
	var state bool = false

	for i:=0; i<H-1; i++{
		for j:=0; j<W; j++{
			if ((S[i][j] == ".") && (S[i+1][j] == "#")) || ((S[i][j] == "#") && (S[i+1][j] == ".")){
				if state == false {
					edge += 1
					state = true
				}
			} else {
				state = false
			}
		}
		state = false
	}

	for j:=0; j<W-1; j++{
		for i:=0; i<H; i++{
			if ((S[i][j] == ".") && (S[i][j+1] == "#")) || ((S[i][j] == "#") && (S[i][j+1] == ".")){
				if !state {
					edge += 1
					state = true
				}
			} else {
				state = false
			}
		}
		state = false
	}

	fmt.Printf("%d\n", edge)
	
}
