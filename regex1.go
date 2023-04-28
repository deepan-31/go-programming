package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	/*
		reStr := "The ape was at the apex"
		match, _ := regexp.MatchString("(ape[^ ]?", reStr)
		pl(match)
	*/
	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")
	pl("Matching string: ", r.MatchString(reStr2))
	pl("find string: ", r.FindString(reStr2))
	pl("Indexg: ", r.FindStringIndex(reStr2))
	pl("MatchString: ", r.FindAllString(reStr2, -1))
	pl("MatchString Index: ", r.FindAllStringSubmatchIndex(reStr2, -1))
}
