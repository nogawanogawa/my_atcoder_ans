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
	fmt.Scanf("%d",&N)

	var min_p int = -1

	var A, P, X int 
	for i:=0; i<N; i++{
		A = nextInt()
		P = nextInt()
		X = nextInt()

		if (A < X) && (min_p < 0){
			min_p = P
		} else if (A < X) && (P < min_p){
			min_p = P
		}
	}

	fmt.Printf("%d\n", min_p)
}
