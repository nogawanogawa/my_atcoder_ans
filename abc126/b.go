package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	s := strings.Split(S, "")

	first, _ := strconv.Atoi(strings.Join(s[:2], ""))
	second, _ := strconv.Atoi(strings.Join(s[2:], ""))

	var flag_1, flag_2 int
	if (0 < first) && (first < 13) {
		flag_1 = 0
	} else {
		flag_1 = 1
	}

	if (0 < second) && (second < 13) {
		flag_2 = 0
	} else {
		flag_2 = 2
	}

	if flag_1+flag_2 == 0 {
		fmt.Printf("%s\n", "AMBIGUOUS")
	} else if flag_1+flag_2 == 1 {
		fmt.Printf("%s\n", "YYMM")
	} else if flag_1+flag_2 == 2 {
		fmt.Printf("%s\n", "MMYY")
	} else {
		fmt.Printf("%s\n", "NA")
	}

}
