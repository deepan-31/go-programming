package main

import (
	"fmt"
)

var pl = fmt.Println

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}
func main() {

	pl("Generic sum: ", getSumGen(5.01, 4.23))
}
