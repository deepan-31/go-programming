package main

import (
	"fmt"
)

var pl = fmt.Println

type contact struct {
	fname, lname, phone string
}
type buisness struct {
	name, address string
	contact
}

func (b buisness) info() {
	fmt.Printf("Contact at %s is %s %s", b.name, b.contact.fname, b.contact.lname)
}

func main() {
	con1 := contact{
		"James", "Wang", "543-456-567",
	}
	bus1 := buisness{
		"ABC plumbing",
		"234 norhting st",
		con1,
	}
	bus1.info()

}
