package main

import (
	"fmt"
	"strings"
)

func main() {

	var H, W int
	fmt.Scanf("%d %d", &H, &W)

	s := make([][]string, H)
	var S string
	for i := 0; i < H; i++ {
		fmt.Scanf("%s", &S)
		s[i] = strings.Split(S, "")
	}

	var state string
	var ans int = 1

	for i := 0; i < W+H-1; i++ {
		state = "."

		for j := 0; j <= i; j++ {
			if j >= W {
				continue
			}
			if i-j >= H {
				continue
			}

			if s[i-j][j] == "R" {
				if state == "B" {
					fmt.Printf("%d\n", 0)
					return
				}
				state = "R"
			} else if s[i-j][j] == "B" {
				if state == "R" {
					fmt.Printf("%d\n", 0)
					return
				}
				state = "B"
			}
		}

		if state == "." {
			ans = (ans * 2) % 998244353
		}
	}

	fmt.Printf("%d\n", ans)

}
