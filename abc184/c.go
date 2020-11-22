package main

import (
	"fmt"
	"math"
)

func main() {
    var r1, r2, c1, c2 int
    fmt.Scanf("%d %d", &r1, &c1)
    fmt.Scanf("%d %d", &r2, &c2)
	
	var ans int 

	if (r1 == r2) && (c1 == c2){
		ans = 0		
	} else if (r1+c1 == r2+c2) || (r1-c1 == r2-c2) || math.Abs(float64(r1-r2)) + math.Abs(float64(c1-c2)) <=3{
		ans = 1
  	} else if ((r2 - r1 + c2 - c1) % 2 == 0) || (math.Abs(float64(r2 - r1 - c2 + c1)) <= 3) || (math.Abs(float64(r2 - r1 + c2 - c1)) <= 3){
		ans = 2
	} else {
		ans = 3
	}

	fmt.Printf("%d\n", ans)
	
}
