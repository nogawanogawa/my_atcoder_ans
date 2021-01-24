package main
 
import (
  "fmt"
  "strings"
)
 
func dfs(B [][]string, i int, j int){

  B[i][j] = "x"

  var next_i, next_j int

  next_i = i-1
  next_j = j
  if next_i >= 0{
      if B[next_i][next_j] == "o"{
        dfs(B, next_i, next_j)
      }    
  }
  
  next_i = i+1
  next_j = j
  if next_i < 10{
    if B[next_i][next_j] == "o"{
      dfs(B, next_i, next_j)
    }    
  }

  next_i = i
  next_j = j-1
  if next_j >= 0{
    if B[next_i][next_j] == "o"{
      dfs(B, next_i, next_j)
    }    
  }

  next_i = i
  next_j = j+1
  if next_j < 10{
    if B[next_i][next_j] == "o"{
      dfs(B, next_i, next_j)
    }    
  }
  return

}

func main() {

  A := make([][]string, 10)
  B := make([][]string, 10)
  for i := 0; i < 10; i++ {
    A[i] = make([]string, 10)
    B[i] = make([]string, 10)
  }

  var row string
  for i := 0; i < 10; i++ {
    fmt.Scanf("%s", &row)
    rows := strings.Split(row, "")

    for j:= 0; j < 10; j++ {
      A[i][j] = rows[j]     
    }
  }

  var flag bool
  for i := 0; i < 10; i++ {
    for j:= 0; j < 10; j++ {
      if A[i][j] == "x"{
        flag = true
        for k := 0; k < 10; k++ {
          _ = copy(B[k], A[k])        
        }

        dfs(B, i, j)

        // 確認
        for k := 0; k < 10; k++ {
          for l:= 0; l < 10; l++ {
            if B[k][l] == "o"{
              flag = false
              break
            }
          }
          if flag == false{
            break
          }
        }
        if flag == true {
          fmt.Printf("YES\n")	
          return
        }
      }
    }
  }

  fmt.Printf("NO\n")
}
 