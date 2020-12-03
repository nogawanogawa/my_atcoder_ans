package main
import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	var AC, WA, TLE, RE int
	var s string
	for i:=0;i<N;i++ {
		fmt.Scanf("%s", &s)

		if s == "AC" {
			AC++
		} else if s == "WA"{
			WA++
		} else if s == "TLE" {
			TLE++
		} else if s == "RE"{ 
			RE ++ 
		}
	}
	fmt.Printf("AC x %d\n", AC)
	fmt.Printf("WA x %d\n", WA)
	fmt.Printf("TLE x %d\n", TLE)
	fmt.Printf("RE x %d\n", RE)
}