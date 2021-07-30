package main

import (
	"errors"
	"fmt"
)

const (
	ErrNotFound = iota
)

// START OMIT
type CustomError struct {
	Operation string
	Type      int
	Err       error
}

func NewCustomError(op string, t int, err error) error {
	return &CustomError{op, t, err}
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Op: [%s], Type: [%d], Err:[%s]", c.Operation, c.Type, c.Err)
}

// END OMIT

type Ingredient string

func getIngredients() ([]Ingredient, error) {

	tortilla, err := getTortilla()
	if err != nil {
		return nil, fmt.Errorf("could not get tortilla: %w", err)
	}

	meat, err := getMeat()
	if err != nil {
		return nil, fmt.Errorf("could not get meat: %w", err)
	}

	toppings, err := getToppings()
	if err != nil {
		return nil, fmt.Errorf("could not get toppings: %w", err)
	}

	return []Ingredient{tortilla, meat, toppings}, nil
}

func main() {

	ingredients, err := getIngredients()
	if err != nil {
		var custErr *CustomError
		if errors.As(err, &custErr) {
			if custErr.Type == ErrNotFound {
				fmt.Println("Ah, we don't have enough ingredients! CustomErr:", custErr)
				return
			}
			fmt.Println("CustomError: ", custErr)
			return
		}
		panic(err)
	}

	makeTaco(ingredients)
}

func makeTaco(ingredients []Ingredient) {
	fmt.Printf("making taco with %s", ingredients)
}

func getTortilla() (Ingredient, error) {
	return "", NewCustomError("getTortilla", ErrNotFound, errors.New("no such ingredient"))
}

func getMeat() (Ingredient, error) {
	return "meat", nil
}

func getToppings() (Ingredient, error) {
	return "toppings", nil
}
