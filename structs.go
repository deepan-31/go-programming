package main

import (
	"fmt"
)

var pl = fmt.Println

/*
type customer struct {
	name    string
	address string
	bal     float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f", c.name, c.bal)
}
func newCustAdd(c *customer, address string) {
	c.address = address
}
*/

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}
func main() {
	/*
		var tS customer
		tS.name = "Tom Smith"
		tS.address = "S Downing street"
		tS.bal = 234.56
		getCustInfo(tS)
		newCustAdd(&tS, "123 south st")
		pl("address: ", tS.address)
		sS := customer{"Sally smith", "123 main", 0.0}
		pl("Name: ", sS.name)
	*/
	rect1 := rectangle{10.0, 15.0}
	pl("rECT AREA", rect1.Area())

}
