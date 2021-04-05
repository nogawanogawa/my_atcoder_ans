package main

import (
    "fmt"
    "math"
)


func main() {
    var a, b, x float64
    fmt.Scanf("%f %f %f", &a, &b, &x)

    if a * a * b / 2 > x{
        tan_theta := 2 * x / (b * b * a)
        result := math.Atan(tan_theta) * 180 /math.Pi 
        fmt.Printf("%f\n", 90-result)    
    } else {
        c:=2*b - 2*x / (a*a)
        tan_theta := c / a
        result := math.Atan(tan_theta) * 180 /math.Pi 
        fmt.Printf("%f\n", result)  
    }
}
