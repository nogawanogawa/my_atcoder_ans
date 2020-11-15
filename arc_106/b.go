package main
 
import (
  "fmt"
)
 
func main() {
  var N, a int
  fmt.Scan(&N)
  
  ans := 0
  for i := 0; i < N; i++ {
    fmt.Scan(&a)
    ans = gcd(ans ,a)
  }
  fmt.Printf("%d\n", ans)
}
 
func gcd(a, b int) int {
  if b == 0 {
    return a
  }
  return gcd(b, a % b)
}