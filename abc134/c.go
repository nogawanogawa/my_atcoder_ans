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

	A := make([]int, N)
	var first int = 0 
	var second int =0

	for i:=0; i<N; i++{
		A[i] = nextInt()
		if first < A[i]{
			second = first
			first = A[i]
		} else if (first >= A[i]) && (second < A[i]){
			second = A[i]
		}
	}

	for i:=0; i<N; i++{
		if first == A[i]{
			fmt.Printf("%d\n", second)
		} else {
			fmt.Printf("%d\n", first)
		}
	}

}
