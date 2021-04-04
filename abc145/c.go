package main

import (
    "fmt"
    "math"
)

func search(cost float64, now int, array [][]float64, point []int) float64{

    var cost_ float64
    point_ := make([]int, len(point))
    
    if len(point) == 1{
        cost_ = cost + array[now][point[0]]
    } else {
        for i:=0; i<len(point); i++{
            copy(point_, point)
            c := array[now][point[i]]
            cost_ += search(cost+c, point[i], array, append(point_[:i], point_[i+1:]...))
        }
    }

    return cost_
}

func main() {
    var N int
    fmt.Scanf("%d", &N)

    x := make([]int, N)
    y := make([]int, N)

    for i:=0; i<N; i++{
        fmt.Scanf("%d %d", &x[i], &y[i])
    }

    array := make([][]float64, N)
    for i:=0; i<N; i++{
        array[i] = make([]float64, N)
    }

    for i:=0; i<N; i++{
        for j:=0; j<N; j++{
            array[i][j] = math.Sqrt(float64((x[i] - x[j]) * (x[i] - x[j]) + (y[i] - y[j]) * (y[i] - y[j])))
        }
    }

    point := make([]int, N)
    for i:=0; i<N; i++{
        point[i] = i
    }

    var cost float64
    var path float64 = 1
    point_ := make([]int, N)

    for i:=0; i<N; i++{
        copy(point_, point)
        cost += search(0, i, array, append(point_[:i], point_[i+1:]...))
        path *= (float64(i)+1)
    }

    result := cost / path

    fmt.Printf("%f\n", result)
}
