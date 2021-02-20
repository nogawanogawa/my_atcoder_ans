package main

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func calc (a int) int{

	a_s := strconv.Itoa(a)

	a_s_1 := strings.Split(a_s, "")
	a_s_2 := strings.Split(a_s, "")

	sort.Strings(a_s_1)
    sort.Slice(a_s_2, func(i, j int) bool {
        return a_s_2[i] > a_s_2[j]
    })

	a_1 := strings.Join(a_s_1, "")
	a_2 := strings.Join(a_s_2, "")

	a__1, _ := strconv.Atoi(a_1)
	a__2, _ := strconv.Atoi(a_2)
	
	ans := a__2 - a__1
	return ans
}

func main() {
    var N, K int
    fmt.Scanf("%d %d", &N, &K)

	for i:=0; i<K ; i++{
		N = calc(N)
	}

	fmt.Printf("%d\n", N)	
}
