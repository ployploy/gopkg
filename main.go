package main

import (
	"fmt"

	"github.com/ployploy/gopkg/calculator"
)

func main() {
	area := calculator.Area{}
	result := area.SquarArea(3, 4)
	fmt.Printf("%d", result)
}
