package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Point struct {
	i int
	x int
	y int
}

type Pair struct {
	max      int
	min      int
	distance int
}

func main() {

	sc.Split(bufio.ScanWords)

	var N int
	fmt.Scanf("%d", &N)

	point := make([]Point, N)

	pair := make([]Pair, 4)

	for i := 0; i < N; i++ {
		point[i].x = nextInt()
		point[i].y = nextInt()
		point[i].i = i
	}

	sort.Slice(point, func(i, j int) bool {
		return point[i].x < point[j].x
	})

	pair[0].distance = point[N-1].x - point[0].x
	pair[0].max = point[N-1].i
	pair[0].min = point[0].i

	if (point[N-2].x - point[0].x) > (point[N-1].x - point[1].x) {
		pair[1].distance = point[N-2].x - point[0].x
		pair[1].max = point[N-2].i
		pair[1].min = point[0].i
	} else {
		pair[1].distance = point[N-1].x - point[1].x
		pair[1].max = point[N-1].i
		pair[1].min = point[1].i
	}

	sort.Slice(point, func(i, j int) bool {
		return point[i].y < point[j].y
	})

	pair[2].distance = point[N-1].y - point[0].y
	pair[2].max = point[N-1].i
	pair[2].min = point[0].i

	if (point[N-2].y - point[0].y) > (point[N-1].y - point[1].y) {
		pair[3].distance = point[N-2].y - point[0].y
		pair[3].max = point[N-2].i
		pair[3].min = point[0].i
	} else {
		pair[3].distance = point[N-1].y - point[1].y
		pair[3].max = point[N-1].i
		pair[3].min = point[1].i
	}

	sort.Slice(pair, func(i, j int) bool {
		return pair[i].distance > pair[j].distance
	})

	if ((pair[0].max == pair[1].max) && (pair[0].min == pair[1].min)) || ((pair[0].max == pair[1].min) && (pair[0].min == pair[1].max)) {
		fmt.Printf("%d\n", pair[2].distance)
	} else {
		fmt.Printf("%d\n", pair[1].distance)
	}

}
