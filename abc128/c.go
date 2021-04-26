package main

import (
	"fmt"
)

func main() {

	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	matrix := make([][]int, M)
	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
	}

	var k, s int
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &k)
		for j := 0; j < k; j++ {
			fmt.Scanf("%d", &s)
			matrix[i][s-1] = 1
		}
	}

	p := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &p[i])
	}

	light := make([]int, M)
	var count int
	var flag bool

	// bitsがn個の要素の各パターンを表す
	for bits := 0; bits < (1 << uint64(N)); bits++ {
		for j := 0; j < M; j++ {
			light[j] = 0
		}

		// bitsの各要素についてのループ
		for i := 0; i < N; i++ {
			// bitsのi個目の要素の状態がonかどうかチェック
			if (bits>>uint64(i))&1 == 1 {
				for j := 0; j < M; j++ {
					light[j] += matrix[j][i]
				}
			}
		}

		flag = true
		for j := 0; j < M; j++ {
			if light[j]%2 != p[j] {
				flag = false
				break
			}
		}

		if flag {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}
