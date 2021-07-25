package main

import (
	"fmt"
	"sort"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {

	var N int
	fmt.Scanf("%d", &N)

	a := make([]int, N)
	b := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &b[i])
	}

	m := make(map[int]int)

	var v int
	for j := 0; j < N; j++ {
		v = a[0] ^ b[j]
		_, ok := m[v]
		if ok == false {
			m[v] = 1
		}
	}

	for i := 1; i < N; i++ {
		for j := 0; j < N; j++ {
			v = a[i] ^ b[j]

			_, ok := m[v]
			if ok == true {
				if m[v] >= i {
					m[v] = i + 1
				} else {
					delete(m, v)
				}
			} else {
				delete(m, v)
			}
		}
	}

	var l []int

	for k := range m {
		if m[k] != N {
			delete(m, k)
		} else {
			l = append(l, k)
		}
	}

	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})

	fmt.Printf("%d\n", len(m))
	for i := 0; i < len(l); i++ {
		fmt.Printf("%d\n", l[i])
	}
}
