package main

import (
	"fmt"
)

var pl = fmt.Println

func fact(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * fact(num-1)
	}
}
func main() {
	pl(fact(4))

}
