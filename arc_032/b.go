package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type UnionFind struct {
	par []int64
}

func newUnionFind(N int64) *UnionFind {
	u := new(UnionFind)
	
	u.par = make([]int64, N+1)

	var i int64
    for i=0; i<N+1; i++ {
		u.par[i] = -1
	}
	
    return u
}

// 根を返す
func (u UnionFind) root(x int64) int64 {
    if u.par[x] < 0 {
        return x // 初期値マイナスのため自分が根
    }
    u.par[x] = u.root(u.par[x]) // 親をたどっていって根を返す
    return u.par[x]
}

// union クエリ
func (u UnionFind) unite(x, y int64) {
    xr := u.root(x)
    yr := u.root(y)
    if xr == yr {
        return
    }
    u.par[yr] += u.par[xr]
    u.par[xr] = yr
}

// find クエリ
func (u UnionFind) same(x, y int64) bool {
    if u.root(x) == u.root(y) {
        return true
    }
    return false
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int64 {
    sc.Scan()
    i, e := strconv.ParseInt(sc.Text(),10,32)
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
	sc.Split(bufio.ScanWords)
	
	var N, M int64
	fmt.Scanf("%d %d",&N, &M)

	var i int64
	
	u := newUnionFind(N)

	var a, b int64

	for i=0; i<M; i++{
		a = nextInt()	
		b = nextInt()

		u.unite(a,b)
	}

	dict := make(map[int64]int64)
	var count int = 0
	for i=1; i<N+1; i++{

		_, ok := dict[u.root(i)]
		if ok==false{
			dict[u.root(i)] = i
			count ++
		}
	}

	fmt.Printf("%d\n", count-1)
}
