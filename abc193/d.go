package main

import (
    "fmt"
    "strings"
    "strconv"
)

func score(s []int) int{
    hand := make(map[int]int)
    for i:=0; i<5; i++{
        _, isThere := hand[s[i]]
        if !isThere {
            hand[s[i]] = 1
        } else {
            hand[s[i]] = hand[s[i]] + 1
        }         
    }

    var sc int = 0
    for i:=1; i<10; i++{
        _, isThere := hand[i]
        var r int = 1
        if isThere {
            for j:=0; j<hand[i]; j++{
                r = r * 10
            }
            sc += i * r
        }else {
            sc += i
        }
    }
     
    return sc
}

func prob(m map[int]int,K int, s int , t int) float64 {

    var p float64
    if s == t {
        bunshi := m[s]
        bunshi *= m[s] -1
        bunbo := (9*K-8) * (9*K-9) 
        p = float64(bunshi) / float64(bunbo) 
    } else {
        bunshi := m[s]
        bunshi *= m[t]
        bunbo := (9*K-8) * (9*K-9) 
        p = float64(bunshi) / float64(bunbo)
    }

    return p
}

func main() {
    var K int
    var S, T string

    fmt.Scanf("%d", &K)
    fmt.Scanf("%s", &S)
    fmt.Scanf("%s", &T)
    
    m := make(map[int]int)
    for i:=1; i<10; i++{
        m[i] = K
    }

    S_fact := strings.Split(S, "")
    T_fact := strings.Split(T, "")

    for i:=0; i<4; i++{
        n, _ := strconv.Atoi(S_fact[i])
        m[n] = m[n] - 1

        l, _ := strconv.Atoi(T_fact[i])
        m[l] = m[l] - 1
    }

    var result float64 = 0
    for i:=1; i<10; i++{
        for j:=1; j<10; j++{
            p := prob(m, K, i, j)
            var s []int
            for k:=0; k<4; k++{
                tmp, _ := strconv.Atoi(S_fact[k])
                s = append(s, tmp)
            }
            s = append(s, i)
            score_s := score(s)

            var t []int
            for k:=0; k<4; k++{
                tmp, _ := strconv.Atoi(T_fact[k])
                t = append(t, tmp)
            }
            t = append(t, j)
            score_t := score(t)

            if score_s > score_t {
                result += p
            }
        }
    }

	fmt.Printf("%f\n", result)
}
