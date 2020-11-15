package main

import (
    "fmt"
)

func main() {
    var x, y, a, b uint64
    fmt.Scanf("%d %d %d %d", &x, &y, &a, &b)

	var cnt uint64 = 0

	for x < y/a & b/(a-1) {
		cnt++
		x *= a
	}

	cnt += (y-1-x)/b

    fmt.Printf("%d\n", cnt)
}
