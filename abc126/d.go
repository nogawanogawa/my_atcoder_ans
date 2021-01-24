package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type Node struct {
	next []int64
	edge []int64
	color int
}

func dfs(x []Node, j int64){

	nexts := x[j].next
	edges := x[j].edge
	
	for i:=0; i<len(nexts); i++{
		next := nexts[i]
		edge := edges[i]

		if x[next].color == -1 {
			if edge % 2 == 0{
				if x[j].color == 0{
					x[next].color = 0
				} else {
					x[next].color = 1
				}
			} else {
				if x[j].color == 0{
					x[next].color = 1
				} else {
					x[next].color = 0
				}
			}
			dfs(x, next)
		}
	}
}

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
    fmt.Scanf("%d", &N)

	u := make([]int64, N-1)
	v := make([]int64, N-1)
	w := make([]int64, N-1)

	var i int64
	for i=0; i<N-1; i++{
		u[i] = nextInt()
        v[i] = nextInt()
        w[i] = nextInt()
	}

	x := make([]Node, N+1)
	for i=2; i<N+1; i++{
		x[i].color = -1
	}
	for i=0; i<N-1; i++{
		x[u[i]].next = append(x[u[i]].next, v[i])
		x[u[i]].edge = append(x[u[i]].edge, w[i])

		x[v[i]].next = append(x[v[i]].next, u[i])
		x[v[i]].edge = append(x[v[i]].edge, w[i])
	}

	dfs(x, 1)

	for i=1; i<N+1; i++{
		fmt.Printf("%d\n", x[i].color)
	}
}