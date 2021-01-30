package main

import (
    "fmt"
    "math"
)

func main() {
    var N int64
    fmt.Scanf("%d", &N)
    
    max := int64(math.Sqrt(float64(2*N))) 
    a := make([]int64, max+1)

    var i, j int64
    var count int64 = 0
    for i=1; i*i<=2*N; i++{
        if 2*N % i == 0{
            a[j] = i
            j ++
        }
    }

    for i=0; i<max+1; i++{
        if a[i] == 0{
            break
        }
        if  (a[i] % 2) != ((2*N/a[i]) % 2) {
            count += 2
        }
    }

	fmt.Printf("%d\n", count)
}
