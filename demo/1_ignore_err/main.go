package main

import (
	"errors"
	"fmt"
)

// START OMIT

type Ingredient string

func getIngredients() ([]Ingredient, error) {

	tortilla, _ := getTortilla()

	meat, _ := getMeat()

	toppings, _ := getToppings()

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
	return "", errors.New("No such ingredient")
}

func getMeat() (Ingredient, error) {
	return "meat", nil
}

func getToppings() (Ingredient, error) {
	return "toppings", nil
}
