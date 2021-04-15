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

	A := make([]int, N+1)
	B := make([]int, N)
	
	for i:=0; i<N+1; i++{
		A[i] = nextInt()
	}

	for i:=0; i<N; i++{
		B[i] = nextInt()
	}

	var count int = 0

	for i:=0; i<N; i++{
		if A[i] >= B[i] {
			A[i] = A[i] - B[i]
			count += B[i]
			B[i] = 0
		} else {
			count += A[i]
			B[i] = B[i] - A[i]
			A[i] = 0
			
			if A[i+1] >= B[i] {
				A[i+1] = A[i+1] - B[i]
				count += B[i]
				B[i] = 0
			} else {
				count += A[i+1]
				B[i] = B[i] - A[i+1]
				A[i+1] = 0				
			}
		}
	}	
	
	fmt.Printf("%d\n", count)
	

}
