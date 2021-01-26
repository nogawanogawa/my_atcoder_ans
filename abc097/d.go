package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type UnionFind struct {
	par []int
	val []int
}

func newUnionFind(N int, p []int) *UnionFind {
	u := new(UnionFind)
	
	u.par = make([]int, N+1)
	u.val = make([]int, N+1)
	
    for i:=0; i<len(p); i++ {
		u.val[i] = p[i]
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
	
	var N, M int
	fmt.Scanf("%d %d",&N, &M)

	dict := make(map[int]int)

	p := make([]int, N+1)
	for i:=1; i<N+1; i++{
		p[i] = nextInt()
		dict[p[i]] = i
	}
	
	u := newUnionFind(N, p)

	var x, y int

	for i:=0; i<M; i++{
		x = nextInt()	
		y = nextInt()

		u.unite(x,y)
	}

	var count int = 0
	for i:=1; i<N+1; i++{
		if u.root(i) == u.root(dict[i]){
			count ++ 
		}
	}

	fmt.Printf("%d\n", count)
}
