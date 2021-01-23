package main

import (
	"fmt"
)

func dfs(S []string, N int, i int) int64{

    var count int64 = 0

    if i == 0{
        return 1
    }

    count += dfs(S, N, i-1)

    if S[i] == "OR"{

        var c int64 = 1
        for j:=0; j<i; j++{
            c = c * 2
        } 
        count += c
    }

    return count
}

func main() {
    var N int
    fmt.Scanf("%d", &N)
	
	S := make([]string, N+1) // AND/OR
	for i:=1; i<=N; i++{
		fmt.Scanf("%s", &S[i])
	}

    var count int64 = 0

    count += dfs(S, N, N)

	fmt.Printf("%d\n", count)
}
