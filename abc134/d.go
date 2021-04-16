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

	a := make([]int, N)
	for i:=0; i<N; i++{
		a[i] = nextInt()
	}

	b := make([]int, N)
	var sum int = 0

	for i:=N-1; i>=0; i--{
		var count int = 0
		for j:=1; j*(i+1)-1<N; j++{
			count += b[j*(i+1)-1]
		}

		if a[i] == 0{
			if count % 2 == 0{
				b[i] = 0
			} else {
				sum ++
				b[i] = 1
			}
		} else {
			if count % 2 == 0{
				b[i] = 1
				sum ++
			} else {
				b[i] = 0
			}
		}
	}

	fmt.Printf("%d\n", sum)
	for i:=0; i<N; i++{
		if b[i] == 1{
			fmt.Printf("%d ", i+1)
		}
	}
}
