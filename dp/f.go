package main

import (
    "fmt"
    "strings"
)

func main() {

    var s, t string
    fmt.Scanf("%s", &s)
    fmt.Scanf("%s", &t)

    S := strings.Split(s, "")
    T := strings.Split(t, "")

    dp := make([][]int, len(S)+1)
    for i:=0; i<len(S)+1; i++{
        dp[i] = make([]int, len(T)+1)
    }

    for i:=1; i<len(S)+1; i++{
        for j:=1; j<len(T)+1; j++{
            dp[i][j] = dp[i-1][j]
            if dp[i][j] < dp[i][j-1] {
				dp[i][j] = dp[i][j-1]
			}
			if s[i-1] == t[j-1] && dp[i][j] < dp[i-1][j-1]+1 {
				dp[i][j] = dp[i-1][j-1] + 1
			}
        }
    }

    var result string = ""
    i, j := len(S), len(T)
    for (i>0) && (j>0) {
        if (S[i-1] == T[j-1]) && (dp[i-1][j-1]+1 == dp[i][j]){
            result = S[i-1] + result
            i --
            j -- 
            continue
        }
        if dp[i][j-1] == dp[i][j] {
			j--
		} else {
			i--
		}
    }


    fmt.Printf("%s\n", result)
}