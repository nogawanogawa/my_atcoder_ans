package main

import (
	"fmt"
	"strings"
)

func main() {

	var S string
	fmt.Scanf("%s", &S)

	s := strings.Split(S, "")

	Array_c := make([]int, len(s))
	Array_h := make([]int, len(s))
	Array_o := make([]int, len(s))
	Array_k := make([]int, len(s))
	Array_u := make([]int, len(s))
	Array_d := make([]int, len(s))
	Array_a := make([]int, len(s))

	var count int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == "c" {
			count += 1
		}
		Array_c[i] = count % 1000000007
	}

	count = 0
	for i := 0; i < len(s); i++ {
		if s[i] == "h" {
			count += Array_c[i]
		}
		Array_h[i] = count % 1000000007
	}

	count = 0
	for i := 0; i < len(s); i++ {
		if s[i] == "o" {
			count += Array_h[i]
		}
		Array_o[i] = count % 1000000007
	}

	count = 0
	for i := 0; i < len(s); i++ {
		if s[i] == "k" {
			count += Array_o[i]
		}
		Array_k[i] = count % 1000000007
	}

	count = 0
	for i := 0; i < len(s); i++ {
		if s[i] == "u" {
			count += Array_k[i]
		}
		Array_u[i] = count % 1000000007
	}

	count = 0
	for i := 0; i < len(s); i++ {
		if s[i] == "d" {
			count += Array_u[i]
		}
		Array_d[i] = count % 1000000007
	}

	count = 0
	for i := 0; i < len(s); i++ {
		if s[i] == "a" {
			count += Array_d[i]
		}
		Array_a[i] = count % 1000000007
	}

	var result int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == "i" {
			result = (result + (Array_a[i] % 1000000007)) % 1000000007
		}
	}

	fmt.Printf("%d\n", result)

}
