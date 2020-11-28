package main

import (
	"fmt"
	"strings"
)

func main() {
    var n, k int64
    fmt.Scanf("%d %d", &n, &k)

	var s string
    fmt.Scanf("%s", &s)

	S := strings.Split(s, "")

	ir := make([][]string, k)
	var i, j int64
	for i=0; i < k; i++ {
		ir[i] = make([]string, n)
	}

	var index int64 = 1
	var right int64
	for j=0; j<n; j++{
		right = (j + index) % n
		ir[0][j] = judge(S[j], S[right])
	}

	for i=1; i<k; i++{
		index = (index + index) % n
		for j=0; j<n; j++{
			right = (j+index)%n
			ir[i][j] = judge(ir[i-1][j], ir[i-1][right])
		}
	}

	fmt.Printf("%s", ir[k-1][0])
}

func judge(left string, right string) string {
	if left == right {
		return left
	} else if (left =="R") && (right == "P") {
		return "P"
	} else if (left =="R") && (right == "S") {
		return "R"
	} else if (left =="S") && (right == "R") {
		return "R"
	} else if (left =="S") && (right == "P") {
		return "S"
	} else if (left =="P") && (right == "R") {
		return "P"
	} else if (left =="P") && (right == "S") {
		return "S"
	}
	return "S"
}
