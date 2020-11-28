package main
 
import (
  "fmt"
)
 
func main() {
  var n int64
  fmt.Scanf("%d", &n)

  var X int64 = n+1
  var dis int64 =0
  var i int64
  for i=1; ; i++{
    X = X - i
    
    if X < 0{
      break
    }else{
      dis +=1
    }
  }

  ans := n + 1 - dis
  
  fmt.Printf("%d\n", ans)
}
 
