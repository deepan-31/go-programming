package main

import (
	"errors"
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

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDaY(day int) error {

	if (day < 1) || (day > 31) {
		return errors.New("Incorrect day")
	}
	d.day = day
	return nil

}
