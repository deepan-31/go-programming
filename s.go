package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func daysInMonth(m int, y int) time.Time {
	return time.Date(y, time.Month(m), 0, 0, 0, 0, 0, time.UTC)
}
func main() {
	month, year := 1, 1
	fmt.Scanf("%d %d", &month, &year)
	getnoofday := daysInMonth(month, year)
	pl(getnoofday.Day())
}
