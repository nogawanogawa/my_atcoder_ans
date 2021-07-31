package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var X string
	fmt.Scanf("%s", &X)

	x := strings.Split(X, "")

	var cond1 bool = true
	for i := 0; i < len(x); i++ {
		if x[0] != x[i] {
			cond1 = false
		}
	}

	if cond1 == true {
		fmt.Printf("Weak\n")
		return
	}

	var cond2 bool = true
	for i := 0; i < len(x)-1; i++ {
		a, _ := strconv.Atoi(x[i])
		b, _ := strconv.Atoi(x[i+1])
		if (a+1 == b) || ((a == 9) && (b == 0)) {
			cond2 = true
		} else {
			cond2 = false
			break
		}
	}

	if cond2 == true {
		fmt.Printf("Weak\n")
		return
	}
	fmt.Printf("Strong\n")

}
