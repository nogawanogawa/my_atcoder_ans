package main

import (
	"fmt"
	"strings"
)

func search_col(W int, K int, index int, k int, c [][]string, col []int) int{

	var count int
	if index < W-1 {
		count += search_col(W, K, index+1, k, c, col)
		if k-col[index] >= K{
			count += search_col(W, K, index+1, k-col[index], c, col)
		} 		
	} else if index == W-1{
		if  k == K{
			count += 1
		}
		if k-col[index] == K {
			count += 1
		}
	}

	return count
}

func search_row(H int, W int, K int, index int, k int, c [][]string, row []int, col []int) int{
	var count int
	
	new_col := make([]int, W)
	copy(new_col, col)
	
	if index < H-1 {
		count += search_row(H, W, K, index+1, k, c, row, new_col)
		if k-row[index] >= K{
			for i:=0; i<W; i ++{
				if c[index][i] == "#"{
					new_col[i] --
				}
			}
			count += search_row(H, W, K, index+1, k-row[index], c, row, new_col)
		} 
	} else if index == H-1{
		count += search_col(W, K, 0, k, c, new_col)
		if k-row[index] >= K{
			for i:=0; i<W; i ++{
				if c[index][i] == "#"{
					new_col[i] --
				}
			}
			count += search_col(W, K, 0, k-row[index], c, new_col)
		} 
	}

	
	return count
}

func main(){
	var H, W, K int
	fmt.Scanf("%d %d %d",&H, &W, &K)

	c := make([][]string, H)
	var cs string
	for i := 0; i < H; i++ {
		c[i] = make([]string, W)

		fmt.Scanf("%s", &cs)
		c[i] = strings.Split(cs, "")
	}

	var k int 

	row := make([]int, H)
	for i := 0; i < H; i++ {
		for j := 0; j<W; j++{
			if c[i][j] == "#"{
				row[i] ++
				k ++
			}
		}
	}

	col := make([]int, W)
	for i := 0; i < W; i++ {
		for j := 0; j<H; j++{
			if c[j][i] == "#"{
				col[i] ++
			}
		}
	}

	count := search_row(H, W, K, 0, k, c, row, col)

	fmt.Printf("%d\n", count)
}
