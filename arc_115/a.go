package main

import (
    "fmt"
    "strings"
)


func main() {
    var N, M int
    fmt.Scanf("%d %d", &N, &M)

    m := make(map[int]int)
    var s string
    var count int 
    
    for i:=0; i<N; i++{
        count = 0

        fmt.Scanf("%s", &s)  
        s_ := strings.Split(s, "")
        
        for j:=0; j<M; j++{
            if s_[j] == "1"{
                count = count+1
            }
        }

		_, isThere := m[count]

		if !isThere {
			m[count] = 1
		} else {
			m[count] = m[count] + 1
        }
    }

    var result int
    var odd, even int
    for k := range m {
        if k % 2 == 0{
            even += m[k]
        } else {
            odd += m[k]
        }
    }

    result = odd * even

    fmt.Printf("%d\n", result)
}
