package main

import (
	"fmt"
	"sort"
)

type Restaurant struct {
	number int
	city   string
	point  int
}

func main() {
	var N int
	fmt.Scanf("%d", &N)

	restaurant := make([]Restaurant, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%s %d", &restaurant[i].city, &restaurant[i].point)
		restaurant[i].number = i + 1
	}

	sort.Slice(restaurant, func(i, j int) bool {
		if restaurant[i].city < restaurant[j].city {
			return true
		} else if restaurant[i].city == restaurant[j].city {
			if restaurant[i].point > restaurant[j].point {
				return true
			}
		}
		return false
	})

	for i := 0; i < N; i++ {
		fmt.Printf("%d\n", restaurant[i].number)
	}

}
