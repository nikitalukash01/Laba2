package calc

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTockenize(t *testing.T){
	got := tockenize("+12")
	var buffer  []Tocken
	var a1 = Tocken{"+",1}
	var a2 = Tocken{"1",0}
	var a3  = Tocken{"2",0}
	buffer = append(buffer,a1,a2,a3)
	result:= assert.Equal(t,got,buffer)
	fmt.Println(result)
}



