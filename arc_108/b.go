package main
 
import (
  "fmt"
  "strings"
)
 
func check(queue []string) []string{
  if len(queue) >= 3{
    l := len(queue)
    if (queue[l-3] == "f") && (queue[l-2] == "o") && (queue[l-1] == "x"){
      queue = queue[:l-3]
      queue = check(queue)
    } 
  }
  return queue
}

func main() {
  var N int
  fmt.Scanf("%d", &N)
  var s string
  fmt.Scanf("%s", &s)

  c := strings.Split(s, "")

  var queue []string

  for i:=0; i<N; i++{
    queue = append(queue, c[i])
    queue = check(queue)
  }
  fmt.Printf("%d\n", len(queue))

}
 
