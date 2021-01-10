package main

import (
    "bufio"
    "fmt"
    "sort"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int64 {
    sc.Scan()
    i, e := strconv.ParseInt(sc.Text(),10,64)
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
    sc.Split(bufio.ScanWords)
    var N int64
    var C int64
    N = nextInt()
    C = nextInt()

    a := make([]int64, N)
    b := make([]int64, N)
    c := make([]int64, N)

    var amin int64 = 1000000000
    var bmax int64 = 1
    var i int64
    for i=0; i<N; i++ {

        a[i] = nextInt()
        b[i] = nextInt()
        c[i] = nextInt()

        if amin > a[i]{
            amin = a[i]
        }
        if bmax < b[i]{
            bmax = b[i]
        }
    }

    index := make(map[int64]int64)
    for i=0; i<N; i++{
        index[a[i]] = index[a[i]] + c[i]
        index[b[i]+1] = index[b[i]+1] - c[i] 
    }

    keys := make([]int64, 0, len(index))
    for k := range index {
		keys = append(keys, k)
    }
    sort.Slice(keys, func(s, t int) bool {
        return keys[s] < keys[t]
    })

    var cost_sum int64 = 0
    var sum int64 = 0
    var before int64 = 0
    for i:=0; i<len(index); i++{
        // 前まで足し込む
        if cost_sum > C{
            sum += C * (keys[i] - before)
        } else {
            sum += cost_sum * (keys[i] - before)
        }

        // tugiwo keisann 
        cost_sum = cost_sum + index[keys[i]]
        before = keys[i]
    }

    fmt.Printf("%d\n", sum)    
}


