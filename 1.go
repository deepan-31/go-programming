package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	/*
			fmt.Printf("%s %d %c %f %t %o %x\n",
				"stuff", 1, 'A', 3.14, true, 1, 1)
			fmt.Printf("%9f\n", 3.14)
			fmt.Printf("%9.f\n", 3.14)

		for x := 1; x <= 5; x++ {
			fmt.Println(x)
		}
	*/
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print(" Guess a number btw 0 and 50 ^_^")
		fmt.Println("\nrandNum: ", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess == randNum {
			fmt.Printf("\nCorrect Guess!")
			break
		} else {
			fmt.Printf("\nincorrect Guess\n")
		}
	}
}
