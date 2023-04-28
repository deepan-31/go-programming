package main

import (
	"fmt"
)

var pl = fmt.Println

type Animal interface {
	AngrySound()
	HappySound()
}
type Cat string

func (c Cat) Attack() {
	pl("Cat attacks it prey")
}
func (c Cat) Name() string {
	return string(c)
}
func (c Cat) AngrySound() {
	pl("^^^^^^hissss^^^^^^")
}
func (c Cat) HappySound() {
	pl("------purrr-------")
}
func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("catts name", kitty2.Name())
}
