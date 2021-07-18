package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

type Point struct {
	x int64
	y int64
}

type Best struct {
	p1   Point
	p2   Point
	cost int64
}

func calc_cost(A [][]int64, p Point, i int64, j int64, C int64) int64 {
	var cost int64
	var x_, y_ int64

	cost = A[p.y][p.x] + A[i][j]

	if p.y > i {
		y_ = p.y - i
	} else {
		y_ = i - p.y
	}

	if p.x > j {
		x_ = p.x - j
	} else {
		x_ = j - p.x
	}

	cost = cost + C*(x_+y_)
	return cost
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {

	sc.Split(bufio.ScanWords)

	var H, W, C int64
	fmt.Scanf("%d %d %d", &H, &W, &C)

	A := make([][]int64, H)
	var i, j int64
	for i = 0; i < H; i++ {
		A[i] = make([]int64, W)
	}

	for i = 0; i < H; i++ {
		for j = 0; j < W; j++ {
			A[i][j] = nextInt()
		}
	}

	var cost int64

	D := make([][]int64, H)
	for i = 0; i < H; i++ {
		D[i] = make([]int64, W)
	}

	E := make([][]int64, H)
	for i = 0; i < H; i++ {
		E[i] = make([]int64, W)
	}

	for i = 0; i < H; i++ {
		for j = 0; j < W; j++ {
			cost = A[i][j] - C*(i+j)
			if i > 0 {
				cost = min(cost, D[i-1][j])
			}
			if j > 0 {
				cost = min(cost, D[i][j-1])
			}
			D[i][j] = cost
		}
	}

	var goal_cost, start_cost int64
	var min_cost int64 = 1000000000000000000
	for i = 0; i < H; i++ {
		for j = 0; j < W; j++ {

			if (i == 0) && (j == 0) {
				continue
			} else {
				goal_cost = A[i][j] + C*(i+j)
				if i == 0 {
					start_cost = D[i][j-1]
				} else if j == 0 {
					start_cost = D[i-1][j]
				} else {
					start_cost = min(D[i-1][j], D[i][j-1])
				}

				cost = goal_cost + start_cost
				if min_cost > cost {
					min_cost = cost
				}
			}
		}
	}

	for i = 0; i < H; i++ {
		for j = W - 1; j >= 0; j-- {
			cost = A[i][j] - C*(i-j)
			if i > 0 {
				cost = min(cost, E[i-1][j])
			}
			if j < W-1 {
				cost = min(cost, E[i][j+1])
			}
			E[i][j] = cost
		}
	}

	for i = 0; i < H; i++ {
		for j = W - 1; j >= 0; j-- {
			if (i == 0) && (j == W-1) {
				continue
			} else {
				goal_cost = A[i][j] + C*(i-j)
				if i == 0 {
					start_cost = E[i][j+1]
				} else if j == W-1 {
					start_cost = E[i-1][j]
				} else {
					start_cost = min(E[i-1][j], E[i][j+1])
				}

				cost = goal_cost + start_cost
				if min_cost > cost {
					min_cost = cost
				}
			}
		}
	}

	fmt.Printf("%d\n", min_cost)
}
