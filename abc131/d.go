package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 32)
	if e != nil {
		panic(e)
	}
	return int(i)
}

type Task struct {
	A     int
	B     int
	start int
}

func main() {

	sc.Split(bufio.ScanWords)

	var N int
	fmt.Scanf("%d", &N)

	task := make([]Task, N)

	for i := 0; i < N; i++ {
		task[i].A = nextInt()
		task[i].B = nextInt()
		task[i].start = task[i].B - task[i].A
	}

	sort.Slice(task, func(i, j int) bool {
		if task[i].B < task[j].B {
			return true
		} else if task[i].B == task[j].B {
			if task[i].A < task[j].A {
				return true
			}
		}
		return false
	})

	var time int = 0
	var flag bool = true
	for i := 0; i < N; i++ {

		if time > task[i].start {
			flag = false
			break
		} else {
			time += task[i].A
		}
	}

	if flag {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}

}
