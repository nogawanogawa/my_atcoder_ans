package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func search(nodes []node, n int, parent int) {
	nodes[n].parent = parent
	for i:=0; i<len(nodes[n].next); i++{
		if nodes[n].next[i] != parent{
			search(nodes, nodes[n].next[i], n)
		}
	}
}

func add(nodes []node, num int, op map[int]int, parent_count int){
	nodes[num].count = parent_count + op[num]
	for i:=0; i<len(nodes[num].next); i++{
		if nodes[num].next[i] != nodes[num].parent{
			add(nodes, nodes[num].next[i], op, nodes[num].count)
		}
	}
}

type node struct{
	num int
	count int
	parent int
	next []int
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

	nodes := make([]node, N)
	for i:=0; i<N; i++{
		nodes[i].num = i+1
	}

	var a,b int
	for i:=0; i<N-1; i++{
		a = nextInt()
		b = nextInt()
		nodes[a-1].next = append(nodes[a-1].next, b-1)
		nodes[b-1].next = append(nodes[b-1].next, a-1)
	}

	search(nodes, 0, -1)

	op := make(map[int]int)
	var p1, x int
	for i:=0; i<Q; i++{
		p1 = nextInt()
		x = nextInt()
		_, ok := op[p1-1]
		if ok {
			op[p1-1] = op[p1-1] + x
		} else {
			op[p1-1] =  x
		}
	}

	add(nodes, 0, op, 0)


	for i:=0; i<N; i++{
		if i!=N-1{
			fmt.Printf("%d ", nodes[i].count)
		} else {
			fmt.Printf("%d\n", nodes[i].count)
		}
		
	}

	
}
