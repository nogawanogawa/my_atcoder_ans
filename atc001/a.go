package main

import (
    "fmt"
    "strings"
)

func dfs(c [][]string, d [][]bool, i int, j int, H int, W int){
    
    d[i][j] = true

    var next_i, next_j int

    next_i = i-1
    next_j = j
    if next_i >= 0{
        if (d[next_i][next_j] == false) && (c[next_i][next_j] != "#"){
            dfs(c, d, next_i, next_j, H, W)
        }    
    }
    
    next_i = i+1
    next_j = j
    if next_i < H{
        if (d[next_i][next_j] == false) && (c[next_i][next_j] != "#"){
            dfs(c, d, next_i, next_j, H, W)
        }
    }

    next_i = i
    next_j = j-1
    if next_j >= 0{
        if (d[next_i][next_j] == false) && (c[next_i][next_j] != "#"){
            dfs(c, d, next_i, next_j, H, W)
        }
    }

    next_i = i
    next_j = j+1
    if next_j < W{
        if (d[next_i][next_j] == false) && (c[next_i][next_j] != "#"){
            dfs(c, d, next_i, next_j, H, W)
        }
    }
    return
}

func main() {
    var H, W int
    fmt.Scanf("%d %d", &H, &W)

    d := make([][]bool, H)
    c := make([][]string, H)
    for i:=0; i<H; i++{
        c[i] = make([]string, W)
        d[i] = make([]bool, W)
    }


    var start_i, start_j, goal_i, goal_j int
    var row string 
    for i:=0; i<H; i++{

        fmt.Scanf("%s", &row)
        rows := strings.Split(row, "") 

        for j:=0; j<W; j++{
            c[i][j] = rows[j]
            if c[i][j] == "s"{
                start_i = i
                start_j = j
            }
            if c[i][j] == "g"{
                goal_i = i
                goal_j = j
            }
        }
    }

    dfs(c, d, start_i, start_j, H, W)

    if d[goal_i][goal_j]{
        fmt.Printf("Yes\n")
    } else{
        fmt.Printf("No\n")
    }
}