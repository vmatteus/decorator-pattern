package main

import (
	"fmt"

	"github.com/vmatteus/decorator-pattern/icecream"
)

func main() {
	iceCream := &icecream.BaseChocolateCone{
		ExistingIngrediant: &icecream.FlavourButterscotch{
			ExistingIngrediant: &icecream.FlavourVanilla{},
		},
	}

	fmt.Println("Steps to prepare your ice cream:")
	for index, step := range iceCream.GetPreperationSteps() {
		fmt.Printf("Step %d: %s\n", index+1, step)
	}
	fmt.Printf("Total Cost: $%d\n", iceCream.GetCost())
}
