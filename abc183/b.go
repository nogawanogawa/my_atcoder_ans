package main

import (
	"fmt"
)

func main() {
    var S_x, S_y, G_x, G_y int
    fmt.Scanf("%d %d %d %d", &S_x, &S_y, &G_x, &G_y)
	
	var a float64
	a = float64(G_y + S_y) / float64(G_x - S_x)

	var ans float64
	ans = float64(S_y) / a + float64(S_x)

	fmt.Printf("%f\n", ans)
	
}
