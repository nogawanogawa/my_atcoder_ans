package main

import (
	"bufio"
    "fmt"
    "os"
    "strconv"
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

	H := make([]int, N)
	
	for i:=0; i<N; i++{
		fmt.Scanf("%d", &H[i])
	}

	if N == 1{
		fmt.Printf("Yes\n")
	} else if N == 2{
		if H[0] <= H[1]{
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	} else {
		var flag bool = true
		for i:=N-1; i>=1; i--{
			
			if H[i-1] > H[i]+1 {
				flag = false
				break
			} else if H[i]+1 == H[i-1]{
				H[i-1] = H[i-1] -1
			}
		}

		if flag {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}

}
