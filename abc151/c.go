package main

import (
    "fmt"
)

func main() {
    var N, M int
    fmt.Scanf("%d %d", &N, &M)

    flag := make([]bool, N)
    count := make([]int, N)
    var sum, p int
    var S string

    for i:=0; i<M; i++{
        fmt.Scanf("%d %s", &p, &S)
        if S == "AC"{
            if flag[p-1] == false{
                flag[p-1] = true
                sum ++
            }
        } else {
            if flag[p-1] == false{
                count[p-1] += 1
            }
        }
    }

    var ans int
    for i:=0; i<N; i++{
        if flag[i] == true{
            ans += count[i]
        }
    }


    fmt.Printf("%d %d\n",  sum,ans)
}