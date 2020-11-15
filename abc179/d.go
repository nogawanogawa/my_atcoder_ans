package main

import (
    "fmt"
)

func main() {
    var n, k int
    fmt.Scanf("%d %d", &n, &k)

    var s[] int
	var flag bool = true
	 
	for i := 0; i < k; i++ {
		var l, r int 
		fmt.Scanf("%d %d", &l, &r)
		for ; l<=r; l++{
			for i:=0; i<len(s); i++{
				if l == s[i]{
					flag = false
				  	break
				}
			}
			if flag{
				s = append(s, l)
			}
		}
	}

	var f[200001] int
	f[0] = 1

	for i := 1; i<n; i++{
		var count int = 0
		for j:=0; j<len(s); j++{
			if s[j] <= i{
				count += f[i-s[j]]
			}
		}
		f[i] = count
	}

    fmt.Printf("%d\n", f[n-1]%998244353)
}
