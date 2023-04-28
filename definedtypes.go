package main

import (
	"fmt"
)

var pl = fmt.Println

type Tsp float64
type Tbs float64
type ML float64

func tspToMl(tsp Tsp) ML {
	return ML(tsp * 4.92)
}
func TbToML(tsp Tbs) ML {
	return ML(tsp * 14.79)
}
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs Tbs) ToMLs() ML {
	return ML(tbs * 14.79)
}
func main() {

	ml1 := ML(Tsp(3) * 4.92)
	fmt.Println(ml1)
	ml2 := ML(Tbs(3) * 14.79)
	fmt.Println(ml2)
	pl("2 tsp + 4 tsp = ", Tsp(2)+Tsp(4))
	pl("2 tsp > 4 tsp = ", Tsp(2) > Tsp(4))
	fmt.Printf("3 tsp = %.2f mL\n", Tsp(3).ToMLs())
	fmt.Printf("3 tbs = %.2f mL\n", Tbs(3).ToMLs())

}
