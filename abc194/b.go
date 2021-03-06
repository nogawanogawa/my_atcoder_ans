package main

import (
    "fmt"
)

func main() {
	
	var N int
	fmt.Scanf("%d",&N)

	var A_0 int = 100000
	var A_1 int = 100000
	var B_0 int = 100000
	var B_1 int = 100000
	var a_i, b_i int

	var A, B int
	for i:=0; i<N; i++{
		fmt.Scanf("%d %d",&A, &B)
		if A_0 >= A {
			A_1 = A_0
			A_0 = A
			a_i = i
		}
		if B_0 >= B {
			B_1 = B_0
			B_0 = B
			b_i = i
		}
	}

	if a_i == b_i{
		C := A_0 + B_0
		if (C <= A_1) && (C <= B_1) {
			fmt.Printf("%d\n", C)
		} else {
			if A_1 <= B_1{
				if A_1 <= B_0{
					fmt.Printf("%d\n", B_0)
				} else {
					fmt.Printf("%d\n", A_1)
				}
			} else {
				if A_0 <= B_1{
					fmt.Printf("%d\n", B_1)
				} else {
					fmt.Printf("%d\n", A_0)
				}
			}
		}
	} else {
		if A_0 < B_0{
			fmt.Printf("%d\n", B_0)
		} else {
			fmt.Printf("%d\n", A_0)
		}
	}
}
