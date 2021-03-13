package main

import (
    "fmt"
    "sort"
)

type obj struct {
	W int 
	V int
}

type query struct {
	s int 
	e int
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func solve(objs []obj, boxes []int, query query) int {

    var boxes_ []int
    box_tmp := make([]int, len(boxes))
    copy(box_tmp, boxes)

    if query.e != len(boxes){
        boxes_ = append(box_tmp[:query.s-1], box_tmp[query.e:]...)
    } else {
        boxes_ = box_tmp[:query.s-1]
    }

    sort.Slice(boxes_, func(i, j int) bool {
        return boxes_[i] > boxes_[j]
    })

    var values int = 0

    for i:=0; i<len(objs); i++{
        max_size := 10000000
        max_index := -1
        for j:=0; j<len(boxes_); j++{
            if (boxes_[j] >= objs[i].W){
                if (max_size > boxes_[j]){
                    max_size = boxes_[j]
                    max_index = j
                }    
            } else {
                break
            }
        }
        if max_index != -1{
            boxes_ = append(boxes_[:max_index], boxes_[max_index+1:]...)
            values += objs[i].V
        }
    }

    fmt.Printf("%d\n", values)
    return 0
}

func main() {
    var N, M, Q int
    
    fmt.Scanf("%d %d %d", &N, &M, &Q)
    
	objs := make([]obj, N)
    for i:=0; i<N; i++{
        fmt.Scanf("%d %d", &objs[i].W, &objs[i].V)
    }

    sort.Slice(objs, func(i, j int) bool {
        return objs[i].V > objs[j].V
    })

	boxes := make([]int, M)
    for i:=0; i<M; i++{
        fmt.Scanf("%d", &boxes[i])
    }

    queries := make([]query, Q)
    for i:=0; i<Q; i++{
        fmt.Scanf("%d %d", &queries[i].s, &queries[i].e)
    }

    for i:=0; i<Q; i++{
        solve(objs, boxes, queries[i])
    }
}
