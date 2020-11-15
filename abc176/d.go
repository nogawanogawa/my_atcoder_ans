package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	var h, w int
	fmt.Scanf("%d %d",&h, &w)

	var c_h, c_w int
	fmt.Scanf("%d %d", &c_h, &c_w)

	var d_h, d_w int
	fmt.Scanf("%d %d", &d_h, &d_w)

	r := bufio.NewReader(os.Stdin)
	matrix := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(r, &matrix[i])
	}

	fmt.Print(matrix)
}
