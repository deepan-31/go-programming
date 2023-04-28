package main

import (
	"fmt"
)

var pl = fmt.Println

func printTo15() {
	for i := 0; i < 15; i++ {
		pl("Func 1: ", i)
	}
}
func printTo10() {
	i := 10
	for i > 0 {
		pl("Func 2: ", i)
		i--
	}
}
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}
func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}
func main() {
	/*
		go printTo15()
		go printTo10()

		time.Sleep(2 * time.Second)
		print(time.Second)
	*/
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
}
