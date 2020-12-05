package main

import (
    "fmt"
)

func gcd(a, b int64) int64 {
	if b == 0 {
	  return a
	}
	return gcd(b, a % b)
}

func lcm(a int64, b int64) int64{
    return a * b / gcd(a,b)
}

func main() {
    var N int64
    fmt.Scanf("%d", &N)

    var ans int64 = 1
    var i int64 
    for i=2; i<=N; i++{
        ans = lcm(ans, i)
    }

    fmt.Printf("%d\n", ans+1)
}
