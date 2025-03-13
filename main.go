package main

import (
	"fmt"

	"github.com/vmatteus/decorator-pattern/icecream"
)

func main() {

	base := &icecream.BaseChocolateCone{}

	flavour1 := &icecream.FlavourVanilla{
		ExistingIngrediant: base,
	}

	final := &icecream.FlavourButterScotch{
		ExistingIngrediant: flavour1,
	}

	fmt.Println("Steps to prepare your ice cream:")
	for index, step := range final.GetPreperationSteps() {
		fmt.Printf("Step %d: %s\n", index+1, step)
	}
	fmt.Printf("Total Cost: $%d\n", final.GetCost())
}
