package main

import (
	"fmt"
	"strings"
)

func main() {

	var S string
	fmt.Scanf("%s", &S)

	s := strings.Split(S, "")

	var o, x, q int
	for i := 0; i < len(s); i++ {
		if s[i] == "o" {
			o++
		} else if s[i] == "x" {
			x++
		} else {
			q++
		}
	}

	if o > 4 {
		fmt.Printf("%d\n", 0)
	} else if o == 4 {
		fmt.Printf("%d\n", 24)
	} else if o == 3 {
		//　? use
		a := 24
		ans := a * q

		// no use
		ans += 3 * a / 2
		fmt.Printf("%d\n", ans)
	} else if o == 2 {
		//　? * 2 use
		ans := 12 * (q + q*(q-1))

		//　? * 1 use
		ans += 2 * 12 * q

		// no use
		ans += 14
		fmt.Printf("%d\n", ans)
	} else if o == 1 {
		//　? * 3 use
		ans := 4 * (q + q*(q-1)*3 + q*(q-1)*(q-2))

		//　? * 2 use
		ans += 6 * (q + q*(q-1)) // 6

		//　? * 1 use
		ans += 4 * q // 4

		// no use
		ans += 1
		fmt.Printf("%d\n", ans)
	} else {
		fmt.Printf("%d\n", q*q*q*q)
	}

}
