package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	var m1 map[string]string
	m1 = make(map[string]string)
	m2 := make(map[string]string)
	m1["Batman"] = "Bruce Wayne"
	m1["superman"] = "Clark kent"
	m1["flash"] = "Barry Allen"
	m2["lex luther"] = "lex luthor"

	//su := map[int]string{1: "Krypto",
	//2: "Bat Hound"}
	fmt.Printf("Batman is %v\n", m1["Batman"])
}
