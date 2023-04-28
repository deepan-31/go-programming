package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
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

func (d *Date) SetDay(day int) error {

	if (day < 1) || (day > 31) {
		return errors.New("Incorrect day")
	}
	d.day = day
	return nil

}

func (d *Date) SetMonth(m int) error {

	if (m < 1) || (m > 12) {
		return errors.New("Incorrect month")
	}
	d.day = m
	return nil

}

func (d *Date) SetYear(y int) error {

	if (y < 1875) || (y > time.Now().Year()) {
		return errors.New("Incorrect year")
	}
	d.day = y
	return nil

}

func (d *Date) Day() int {
	return d.day
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Year() int {
	return d.year
}
