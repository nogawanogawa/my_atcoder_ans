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

func diff(l, r int) int {
	if l > r {
		return l - r
	} else {
		return r - l
	}
}

func main() {

	sc.Split(bufio.ScanWords)

	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	A := make([]int, N)
	B := make([]int, M)

	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	for i := 0; i < M; i++ {
		B[i] = nextInt()
	}

	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	var i int = 0
	var j int = 0
	var tmp int
	var min int = 1000000001

	for {
		if (i == len(A)) || (j == len(B)) {
			break
		}

		tmp = diff(A[i], B[j])

		if min > tmp {
			min = tmp
		}

		if i == (len(A) - 1) {
			j += 1
		} else if j == (len(B) - 1) {
			i += 1
		} else {
			if A[i] > B[j] {
				j += 1
			} else {
				i += 1
			}
		}
	}

	fmt.Printf("%d\n", min)

}
