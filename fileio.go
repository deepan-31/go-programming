package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	/*
		f, err := os.Create("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		iPrimeArr := []int{2, 3, 5, 7, 11}
		var sPrimeArr []string
		for _, i := range iPrimeArr {
			sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
		}
		for _, num := range sPrimeArr {
			_, err := f.WriteString(num + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		f, err = os.Open("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scan1 := bufio.NewScanner(f)
		for scan1.Scan() {
			pl(scan1.Text())
		}
		if err := scan1.Err(); err != nil {
			log.Fatal(err)
		}
	*/
	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File doesnt exist")
	} else {
		f, err := os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
