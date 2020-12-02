package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	var a,c,b,x,y int
	fmt.Scan(&a,&b)
	for i:=0;i<a;i++ {
		fmt.Fscan(scanner,&x,&y)
		if b*b >= x*x+y*y {
			c++
		}
	}
	fmt.Print(c)
}