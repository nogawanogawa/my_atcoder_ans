package main

import (
	"fmt"
	"strconv"
	"bufio"
    "os"
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
	var X int
	
	fmt.Scanf("%d %d", &N, &X)

	A := make([]int, N)

	for i:=0; i<N; i++{
		A[i] = nextInt()	
	}

	for i:=0; i<len(A); i++{
		if A[i] != X {
			fmt.Printf("%d", A[i])
			if i != len(A)-1{
				fmt.Printf(" ")
			}
		}
	} 
}
