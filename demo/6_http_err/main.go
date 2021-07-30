package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	ErrNotFound = iota
	ErrInternal
)

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

// START OMIT
func (c *CustomError) IsInChain(t int) bool {
	if c.Type == t {
		return true
	}

	if c.Err == nil {
		return false
	}

	var inner *CustomError
	if errors.As(c.Err, &inner) {
		return inner.IsInChain(t)
	}

	return false
}

// END OMIT

type Ingredient string

func getIngredients() ([]Ingredient, error) {

	tortilla, err := getTortilla()
	if err != nil {
		return nil, NewCustomError("getIngredients", ErrInternal, err)
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

	http.HandleFunc("/taqueria/tacos", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		ingredients, err := getIngredients()
		if err != nil {
			log.Println(err)
			var custErr *CustomError
			if errors.As(err, &custErr) {
				if custErr.IsInChain(ErrNotFound) {
					msg := fmt.Sprintf("Ah, we don't have enough ingredients! CustomErr:%s", custErr.Error())
					rw.WriteHeader(http.StatusNotFound)
					rw.Write([]byte(msg))
					return
				}
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		makeTaco(ingredients)

	})

	if err := http.ListenAndServe("localhost:8080", http.DefaultServeMux); err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

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
