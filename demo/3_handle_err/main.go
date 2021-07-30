package main

import (
	"errors"
	"fmt"
)

type Ingredient string

// START OMIT
func getIngredients() ([]Ingredient, error) {

	tortilla, err := getTortilla()
	if err != nil {
		return nil, fmt.Errorf("could not get tortilla: %s", err)
	}

	meat, err := getMeat()
	if err != nil {
		return nil, fmt.Errorf("could not get meat: %s", err)
	}

	toppings, err := getToppings()
	if err != nil {
		return nil, fmt.Errorf("could not get toppings: %s", err)
	}

	return []Ingredient{tortilla, meat, toppings}, nil
}

func main() {

	ingredients, err := getIngredients()
	if err != nil {
		panic(err)
	}

	makeTaco(ingredients)
}

// END OMIT

func makeTaco(ingredients []Ingredient) {
	fmt.Printf("making taco with %s", ingredients)
}

func getTortilla() (Ingredient, error) {
	return "", errors.New("no such ingredient")
}

func getMeat() (Ingredient, error) {
	return "meat", nil
}

func getToppings() (Ingredient, error) {
	return "toppings", nil
}
