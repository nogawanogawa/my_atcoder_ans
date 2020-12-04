package main
import (
	"fmt"
	"strings"
)

func main() {
	var S, T string
	fmt.Scanf("%s", &S)
	fmt.Scanf("%s", &T)

	s := strings.Split(S, "")
	t := strings.Split(T, "")

	L := len(s)

	var count int = 0
	for i:=0; i<L; i++{
		if s[i] != t[i]{
			count ++
		}
	}

	fmt.Printf("%d\n", count)
}