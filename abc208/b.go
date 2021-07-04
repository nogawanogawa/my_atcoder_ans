package main

import (
	"fmt"
)

func main() {

	var P int
	fmt.Scanf("%d", &P)

	var muls []int
	var mul int = 1
	for i := 1; i <= 10000000; i++ {
		mul = mul * i
		if P < mul {
			break
		} else {
			muls = append(muls, mul)
		}
	}

	var n int = 0
	var cand int
	j := len(muls) - 1
	for {
		if P == 0 {
			break
		}
		cand = muls[j]
		if P >= cand {

			P = P - cand
			n++
		} else {
			j--
		}
	}

	fmt.Printf("%d\n", n)
}
