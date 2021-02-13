package main
 
import (
  "fmt"
)
 
func main() {
  var B, C int64
  fmt.Scanf("%d %d", &B, &C)

  var b_0, b_1, b_2, b_3 int64
  var A0, A1, A2  int64

  b_0 = B - C/2 // 下がり続ける
  b_1 = - B - (C-1)/2 // 一周回って下がり続ける
  b_2 = B + (C-2)/2 // 一周回って下がり続け、最後に戻る
  b_3 = - B + (C-1)/2

  if B == 0{
    A0 = - b_0 + 1
    A1 = 0
    A2 = b_2 - B
  } else if B > 0{
    if (b_0 <= 0){
      A0 = B + B
    }  else {
      A0 = B - b_0 + 1
      A0 += b_3 + B
    }
    A1 = - B - b_1 + 1
    A2 = b_2 - B
  } else {
    if (b_1 <= 0){
      A1 = -B - B
    }  else {
      A1 = -B - b_1 + 1
      A1 += b_2
    } 
    A0 = B - b_0 + 1
    A2 = b_3 + B
  }

  fmt.Printf("%d\n", A0+A1+A2)
}
 
