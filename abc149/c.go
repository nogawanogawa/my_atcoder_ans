package main

import (
    "fmt"
)

func isPrime(X int) bool{
    var flag bool = true
    
    for i:=2; i*i<X; i++{
        if X % i == 0{
            flag = false
            break
        }
    } 

    return flag
}

func main() {
    var X int
    fmt.Scanf("%d", &X)

    var result int 
    for i:=X; i<100004; i++{
        if isPrime(i) {
            result = i 
            break
        }
    }

    fmt.Printf("%d\n", result)    
}