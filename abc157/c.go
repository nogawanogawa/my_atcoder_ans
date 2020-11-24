package main

import (
    "fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	use := make([]bool, N)
	ans := make([]int, N)

	//initialize
	if N > 1 {
		ans[0] = 1
	} else {
		ans[0] = 0
	}

	for i:=1 ; i<N; i++{
		ans[i] = 0
	}

	var s, c int
	var err bool = false

	for i:=0 ; i<M; i++{
		fmt.Scanf("%d %d", &s, &c)

		if use[s-1] && ans[s-1] != c {
			err = true
		}

		if (N > 1) && (s == 1) && (c == 0){
			err = true
		} 
		
		ans[s-1] = c
		use[s-1] = true
	}

	if err {
		fmt.Printf("-1\n")
	} else {
		var res int = 0
		var j int = 1
 
		for i:=0; i<N; i++{
			res += j * ans[N-i-1]
			j = j * 10
		}
		fmt.Printf("%d\n", res)
	}
}