package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type UnionFind struct {
    par []int
}

func newUnionFind(N int) *UnionFind {
    u := new(UnionFind)
    u.par = make([]int, N)
    for i := range u.par {
        u.par[i] = -1
    }
    return u
}

// 根を返す
func (u UnionFind) root(x int) int {
    if u.par[x] < 0 {
        return x // 初期値マイナスのため自分が根
    }
    u.par[x] = u.root(u.par[x]) // 親をたどっていって根を返す
    return u.par[x]
}

// union クエリ
func (u UnionFind) unite(x, y int) {
    xr := u.root(x)
    yr := u.root(y)
    if xr == yr {
        return
    }
    u.par[yr] += u.par[xr]
    u.par[xr] = yr
}

// find クエリ
func (u UnionFind) same(x, y int) bool {
    if u.root(x) == u.root(y) {
        return true
    }
    return false
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.ParseInt(sc.Text(),10,32)
    
    if e != nil {
        panic(e)
    }

    return int(i)
}

func main() {
    sc.Split(bufio.ScanWords)

    var N, Q int
    fmt.Scanf("%d %d", &N, &Q)

    u := newUnionFind(N)

    P := make([]int, Q)
    A := make([]int, Q)
    B := make([]int, Q)

    for i:=0; i<Q; i++{
        P[i] = nextInt()
        A[i] = nextInt()
        B[i] = nextInt()
    }

    for i:=0; i<Q; i++{
        if P[i] == 0{
            u.unite(A[i], B[i])
        } else {
            if u.same(A[i], B[i]){
                fmt.Printf("Yes\n")
            } else {
                fmt.Printf("No\n")
            }
        }
    }
}