package main

import (
	"fmt"
	"strconv"
)

var pl = fmt.Println

var name string = "Derek"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}
