package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return int(i)
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	sc.Split(bufio.ScanWords)

	ih := &intHeap{}
	heap.Init(ih) // 使用する際にはまずInitを呼ぶ

	var Q int
	fmt.Scanf("%d", &Q)

	var q, x int
	var ans []int64
	var state int64
	for i := 0; i < Q; i++ {
		q = nextInt()

		if q == 1 {
			x = nextInt()
			heap.Push(ih, x-int(state)) // heap.Push経由で値を投入
		} else if q == 2 {
			x = nextInt()
			state += int64(x)
		} else {
			v := heap.Pop(ih)
			ans = append(ans, int64(v.(int))+state)
		}
	}

	for j := 0; j < len(ans); j++ {
		fmt.Printf("%d\n", ans[j])
	}
}
