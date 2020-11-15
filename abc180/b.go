package main

import (
	"fmt"
	"math"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)
	
	var t int
	manhattan :=0
	euqulid :=0
	chebishev :=0

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &t)
		
		if t > 0 {
			manhattan += t
		}else{
			manhattan -= t
		}

		euqulid += t*t

		if t > 0 {
			if t > chebishev{
				chebishev = t
			}
		}else{
			if -t > chebishev{
				chebishev = -t				
			}  
		}
	}

	fmt.Printf("%d\n", manhattan)
	
    fmt.Printf("%.10f\n", math.Sqrt(float64(euqulid)))
	
	fmt.Printf("%d\n", chebishev)
	
}
