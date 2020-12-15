package main

import (
	"fmt"
)

func main(){
	var N int
	var S string
	fmt.Scanf("%d",&N)
	fmt.Scanf("%s",&S)

	RGB := make([][]int, N)
	for i:=0; i<N; i++{
		RGB[i] = make([]int, 3)
	}

	if S[N-1] == 'R'{
		RGB[N-1][0] ++
	}else if S[N-1] == 'G'{
		RGB[N-1][1] ++
	}else {
		RGB[N-1][2] ++
	}

	for i:=N-2; i>=0; i--{
		if S[i] == 'R'{
			RGB[i][0] = RGB[i+1][0] + 1
			RGB[i][1] = RGB[i+1][1]
			RGB[i][2] = RGB[i+1][2]
		}else if S[i] == 'G'{
			RGB[i][0] = RGB[i+1][0]
			RGB[i][1] = RGB[i+1][1] + 1
			RGB[i][2] = RGB[i+1][2]
		}else {
			RGB[i][0] = RGB[i+1][0]
			RGB[i][1] = RGB[i+1][1]
			RGB[i][2] = RGB[i+1][2] + 1
		}
	}

	var count, j_i int
	for i:=0; i<N-2; i++{
		for j:=i+1; j<N-1; j++{
			j_i = j - i
			if S[i] != S[j]{
				if S[i] == 'R'{
					if S[j] == 'G'{
						count += RGB[j+1][2]
						if j+j_i < N {
							if S[j+j_i] == 'B'{
								count --
							}
						}
					}else if S[j] == 'B'{
						count += RGB[j+1][1]
						if j+j_i < N {
							if S[j+j_i] == 'G'{
							count --
							}
						}
					}
				} else if S[i] == 'G' {
					if S[j] == 'R'{
						count += RGB[j+1][2]
						if j+j_i < N {
							if S[j+j_i] == 'B'{
								count --
							}
						}
					}else if S[j] == 'B'{
						count += RGB[j+1][0]
						if j+j_i < N {
							if S[j+j_i] == 'R'{
								count --
							}
						}
					}
				} else {
					if S[j] == 'G'{
						count += RGB[j+1][0]
						if j+j_i < N {
							if S[j+j_i] == 'R'{
								count --
							}	
						}
					}else if S[j] == 'R'{
						count += RGB[j+1][1]
						if j+j_i < N {
							if S[j+j_i] == 'G'{
								count --
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("%d\n", count)
}
