package main
 
import (
  "bufio"
  "fmt"
  "sort"
  "os"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int64 {
  sc.Scan()
  i, e := strconv.ParseInt(sc.Text(),10,64)
  if e != nil {
      panic(e)
  }
  return i
}

func main() {
  sc.Split(bufio.ScanWords)

  var N, K int
  fmt.Scanf("%d %d", &N, &K)

  a := make([]int64, N)
  box := make([]int64, K)

  for i:=0; i<N; i++{
    a[i] = nextInt()
  }

  sort.Slice(a, func(i, j int) bool {
    return a[i] < a[j]
  })

  var j int = 0
  var tmp int64 = -1

  for i:=0; i<N; i++{
    if a[i] == tmp{
      if box[j] == a[i]{
        box[j] = a[i]+1
      }
      if j+1 < K{
        j = j+1
      }
    } else {
      if box[0] == a[i]{
        box[0] = a[i]+1
      }
      tmp = a[i]
      if 1 < K{
        j = 1
      }
    }
  }

  var sum int64
  for i:=0; i<K; i++{
    sum += box[i]
  }

  fmt.Printf("%d\n", sum)
}
 

