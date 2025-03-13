package icecream

type FlavourButterScotch struct {
	ExistingIngrediant IIceCreamIngredient
}

func (ingrediant *FlavourButterScotch) GetPreperationSteps() []string {
	step := "1 scoop Butterscotch"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *FlavourButterScotch) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 5 + oldCost
}
