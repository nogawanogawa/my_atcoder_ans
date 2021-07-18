package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
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
	B := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}
	for i := 0; i < N; i++ {
		B[i] = nextInt()
	}
	for i := 0; i < N; i++ {
		C[i] = nextInt()
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})
	sort.Slice(C, func(i, j int) bool {
		return C[i] < C[j]
	})

	var b_i, c_i int
	var count int
	var flag_b, flag_c bool = false, false
	for i := 0; i < N; i++ {
		flag_b = false
		flag_c = false
		a := A[i]
		for j := b_i; j < N; j++ {
			if a < B[j] {
				b := B[j]
				for k := c_i; k < N; k++ {
					if b < C[k] {
						count += 1
						//fmt.Printf("%d %d %d\n", A[i], B[j], C[k])
						flag_b = true
						flag_c = true
						b_i = j + 1
						c_i = k + 1
					}
					if k == N-1 {
						fmt.Printf("%d\n", count)
						return
					}
					if flag_c == true {
						break
					}

				}
				if flag_b == true {
					break
				}
			}
			if j == N-1 {
				fmt.Printf("%d\n", count)
				return
			}
		}
	}

}
