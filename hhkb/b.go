package main

import (
	"fmt"
	"strings"
)

func main() {
    var H, W int
    fmt.Scanf("%d %d", &H, &W)

	S := make([][]string, H)
	for i := 0; i < H; i++ {
		var tmp string
		fmt.Scanf("%s", &tmp)			

		S[i] = strings.Split(tmp, "")
	}

	var count int
	// 横方向
	for i := 0; i < H; i++ {
		for j := 0; j<W-1; j++{
			if S[i][j] == "." && S[i][j+1] == "."{
				count += 1
			}
		}
	}

	// 縦方向
	for j := 0; j<W; j++{
		for i := 0; i < H-1; i++ {
			if S[i][j] == "." && S[i+1][j] == "."{
				count += 1
			}
		}
	}

    fmt.Printf("%d\n", count)
}
