package main

import (
	"fmt"
)

var pl = fmt.Println

func useFunction(f func(int, int) int, x, y int) {
	pl("Answer : ", (f(x, y)))
}
func sumVals(x, y int) int {
	return x + y
}
func main() {
	intSum := func(x, y int) int { return x + y }
	pl("5+4 = ", intSum(5, 4))
	samp1 := 1
	changevar := func() {
		samp1 += 1
	}
	changevar()
	pl("sample = ", samp1)

	useFunction(intSum, 5, 8)
}
