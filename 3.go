package main

import (
	"fmt"
	"reflect"
	"time"
)

var print = fmt.Print

func main() {
	const d int = 24
	print(d, reflect.TypeOf(d), "\n")

	i := 3
	for i > 0 {
		print(i, "\n")
		i--
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		print("It is a weekend\n")
	default:
		print("It is week day\n")
	}
	print(reflect.TypeOf(time.Saturday))

	a := [...]int{1, 2, 3, 4}
	print("\n", a)

	for i, val := range a {
		print(i, ":", val, "\n")
	}
}
