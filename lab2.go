package calc

import (
	"fmt"
)

//Number : adding new constant
const Number = 0

//Operator : adding new constant
const Operator = 1

// Tocken : declarated a new structure
type Tocken struct {
	literal    string
	tockenType int
}

// Cursor : declared a new structure
/*type Cursor struct{
	buf []Tocken
	index int
}
//Next : added new method
func(c Cursor) Next() Tocken{
	index := c.index
	c.index ++
	return c.buf[index]
}*/


func main() {
	fmt.Println(tockenize("+5*-423"))
	buff := tockenize("+5*-423")
	fmt.Println(calc(buff))

}
