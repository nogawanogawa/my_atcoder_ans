package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	dict := make(map[string]int)
	var string_list []string
	var str string
	for i:=0; i<N; i++{
		fmt.Scanf("%s", &str)
		str_slice := strings.Split(str, "")
		sort.Strings(str_slice)
		str = strings.Join(str_slice, "")

		_, ok := dict[str]
		if ok {
			dict[str] = dict[str] + 1
		} else {
			dict[str] = 1
			string_list = append(string_list, str)
		}
		
	}

	var count int64 =0  
	for i:=0; i<len(string_list); i++{
		if dict[string_list[i]] != 0{
			count += int64(dict[string_list[i]] * (dict[string_list[i]]-1) / 2)
		}
	} 

	fmt.Printf("%d\n", count)
}
