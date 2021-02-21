package main
 
import (
  "fmt"
)

func tail_2(A int, B int) int{
  a := A % 100
  b := B % 1000

  var ans int = 1
  for i:=0; i<b; i++{
    ans = (ans * a) % 100
  }

  return ans
}

func tail(A int ,B int) int{
  a := A % 10

  var ans int 
  if (a == 0) || (a == 1) || (a == 5) || (a == 6){
    ans = a
  } else if a == 2 {
    if B % 4 == 0{
      ans = 6
    } else if B % 4 == 1{
      ans = 2      
    } else if B % 4 == 2{
      ans = 4
    } else {
      ans = 8
    }
  } else if a == 3 {
    if B % 4 == 0{
      ans = 1
    } else if B % 4 == 1{
      ans = 3      
    } else if B % 4 == 2{
      ans =9
    } else {
      ans = 7
    }
  } else if a == 4 {
      if B % 2 == 0{
        ans = 6
      } else {
        ans = 4    
      }
  } else if a == 7 {
    if B % 4 == 0{
      ans = 1
    } else if B % 4 == 1{
      ans = 7      
    } else if B % 4 == 2{
      ans = 9
    } else {
      ans = 3
    }
  } else if a == 8 {
    if B % 4 == 0{
      ans = 6
    } else if B % 4 == 1{
      ans = 8      
    } else if B % 4 == 2{
      ans = 4
    } else {
      ans = 2
    }
  } else if a == 9 {
    if B % 2 == 0{
      ans = 1
    } else {
      ans = 9 
    }   
  } 

  return ans

}
 
func main() {
  var A, B, C int
  fmt.Scanf("%d %d %d", &A, &B, &C)

  var ans int 
  BC := tail_2(B, C)  
  ans = tail(A, BC)

  fmt.Printf("%d\n", ans)
}
 
