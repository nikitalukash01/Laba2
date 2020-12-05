package calc

import (
	"strconv"
)
func calc(el []Tocken) int {
	var newStack []int
	for i := len(el) - 1; i >= 0; i-- {
		if el[i].tockenType == 0 {
			ty, _ := strconv.Atoi(el[i].literal)
			newStack = append(newStack, ty)
		} else {
			var a, b int
			a, b, newStack = newStack[len(newStack)-1], newStack[len(newStack)-2], newStack[:len(newStack)-2]
			if el[i].literal == "+" {
				newStack = append(newStack, a+b)
			} else {
				if el[i].literal == "-" {
					newStack = append(newStack, a-b)
				} else {
					if el[i].literal == "*" {
						newStack = append(newStack, a*b)
					} else {
						if el[i].literal == "/" {
							newStack = append(newStack, a/b)
						}
					}
				}

			}
		}

	}
	return newStack[0]
}
