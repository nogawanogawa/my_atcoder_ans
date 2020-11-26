package main

import (
    "fmt"
    "sort"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    dict := make(map[string]int)

    var word string 
    var max int =0

    for i:=0; i<N; i++{

        fmt.Scanf("%s", &word)
        v, ok := dict[word]
        if ok {
            dict[word] = v+1

            if max < v+1 {
                max = v+1
            }
        } else {
            dict[word] = 1

            if max < 1 {
                max = 1
            }
        }
    }

    s := make([]string, 1)

    for k, v := range dict {
        if v == max {
            s = append(s, k)
        }
    }

    sort.Strings(s)

    for i:=0; i<len(s); i++{
        fmt.Printf("%s\n", s[i])
    }
    
}