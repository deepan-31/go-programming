package main

import (
	"fmt"
)

var pl = fmt.Println

func getSum(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You cant divide by zero")
	} else {
		return x / y, nil
	}
}

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

/*
	func changeVal(myPtr *int) {
		*myPtr = 12
	}

	func main() {
		//pl(getQuotient(5, 0))
		//vArr := []int{1, 2, 3, 4}
		//pl(getArraySum(vArr))
		f3 := 5
		pl("f3 before function: ", f3)
		changeVal(&f3)
		pl("f3 after function: ", f3)
	}
*/
func main() {
	f4 := 10
	var f4Ptr *int = &f4
	pl("f4 address: ", f4Ptr)
	pl("f4 Value: ", *f4Ptr)
	*f4Ptr = 11
	pl("f4 Value: ", *f4Ptr)
}
