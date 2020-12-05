package main
 
import (
  "fmt"
)
 
func main() {
  var N int64
  fmt.Scanf("%d", &N)

  var T string
  fmt.Scanf("%s", &T)

  var ans int64
  
  str := "110"
  var n int64
  n = (N+1) / 3 + 1

  var count int64 = 0 
  var i int64

  var S string
  // 文字列候補の作成
  for i=0; i<n; i++{
    S = S + str
  }

  var flag int64
  for i=0; i<3; i++{
    if T == S[i:i+N]{
      count += 1
      flag = i
    }
  }

  if N != 1{
    ans = (10000000000 - (N+flag-1)/3) * count
  } else {
    ans = (10000000000 - (n-1)) * count
  }
  
  fmt.Printf("%d\n", ans)
}
 

