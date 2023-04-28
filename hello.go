package main

import (
	//"bufio"
	"fmt"
	"math"
	//"log"
	//"os"
)

var pl = fmt.Println

func main() {
	/*
		pl("what is your name?")
		reader := bufio.NewReader(os.Stdin)
		name, err := reader.ReadString('\n')
		if err == nil {
			pl("hello ", name)
		} else {
			log.Fatal(err)
		}

		//int,float,bool
		pl(reflect.TypeOf(25))
		pl(reflect.TypeOf(true))
	*/
	/*
		cv3 := "50000000000"
		cv4, err := strconv.Atoi(cv3)
		pl(cv4, err, reflect.TypeOf(cv4))
		cv5 := 120
		cv6 := strconv.Itoa(cv5)
		pl(cv6, reflect.TypeOf(cv6))

		cv7 := "3.14"
		if cv8, err := strconv.ParseFloat(cv7, 64); err == nil {

			pl(cv8, reflect.TypeOf(cv8))
		}
		cv9 := fmt.Sprintf("%.2f", 3.14)
		pl(cv9, reflect.TypeOf(cv9))
	*/
	//conditional <,>,<=,>=,==,!=;
	//logical &&,||,!
	/*
		iAge := 8
		if (iAge >= 1) && (iAge <= 18) {
			pl("important birthday")
		} else if (iAge == 21) || (iAge == 50) {
			pl("important birthday")
		} else if iAge >= 65 {
			pl("important birthday")
		} else {
			pl("Not an important birthday")
		}
		pl("!true = ", !true)
	*/
	/*
		sv1 := "A word"
		replacer := strings.NewReplacer("A", "Another")
		sv2 := replacer.Replace(sv1)
		pl(sv1, "    ", sv2)
		pl("length: ", len(sv2))
		pl("Contains another: ", strings.Contains(sv2, "Another"))
		pl("o index: ", strings.Index(sv2, "o"))
		pl("Replace: ", strings.Replace(sv2, "o", "0", -1))
		sv3 := "\nSomeWords\n" // \t \" \\
		sv3 = strings.TrimSpace(sv3)
		pl("split: ", strings.Split("a-b-c-d", "-"))
		pl("Lower: ", strings.ToLower(sv2))
	*/
	/*
				rStr := "abcdefg"
				//pl(utf8.RuneCountInString(rStr))
				for i, runeVal := range rStr {
					fmt.Printf("%d : %#U : %c \n", i, runeVal, runeVal)
				}

			now := time.Now()
			pl(now.Year(), now.Month(), now.Day())
			pl(now.Hour(), now.Minute(), now.Second())

		//arthemtic operators +,-,*,/,% ,++,--
		seedSecs := time.Now().Unix()
		rand.Seed(seedSecs)
		randNum := rand.Intn(50) + 1
		pl("Randum: ", randNum)
	*/
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians is %.2f degrees", r90, d90)

}
