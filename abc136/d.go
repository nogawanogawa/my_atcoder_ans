package main

import (
    "fmt"
	"strings"
)

func main() {

    var S string
    fmt.Scanf("%s", &S)

	s := strings.Split(S, "")
	result := make([]int, len(s))
	
	var bin int = 0
	for i:=0; i<len(S); i++{
		
		if s[i] == "R"{
			bin ++
		} else {
			h := bin/2
			result[i] += h
			result[i-1] += bin - h
			bin = 0
		}	
	}
	
	for i:=len(S)-1; i>=0; i--{
		if s[i] == "L"{
			bin ++
		} else {
			h := bin/2
			result[i] += h
			result[i+1] += bin - h
			bin = 0
		}
	}

	for i:=0; i<len(S); i++{
		if i==len(S)-1{
			fmt.Printf("%d\n", result[i])
		} else {
			fmt.Printf("%d ", result[i])
		}
	}
}
