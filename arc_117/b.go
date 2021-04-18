package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sort"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.ParseInt(sc.Text(),10,32)
    if e != nil {
        panic(e)
    }
    return int(i)
}


func main() {
    sc.Split(bufio.ScanWords)

    var N int
    fmt.Scanf("%d", &N)

    A := make([]int, N)
    for i:=0; i<N; i++{
        A[i] = nextInt()
    }

    sort.Slice(A, func(i, j int) bool {
      return A[i] < A[j]
    })

    var result int = 1
    for i:=0; i<N; i++{
      if i == 0 {
        result = result * (A[i] + 1) % 1000000007
      } else if A[i] != A[i-1] {
        result = result * (A[i] - A[i-1] + 1) % 1000000007
      }
    }

    fmt.Printf("%d\n", result)
}
