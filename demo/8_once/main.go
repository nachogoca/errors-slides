package main

import (
	"errors"
	"log"
)

// START OMIT
func bar() error {
	return errors.New("boom")
}

func foo() error {
	if err := bar(); err != nil {
		// DON'T DO THIS, HANDLE THE ERROR ONLY ONCE
		log.Printf("error: %s", err.Error())
		return err
	}
	return nil
}

func main() {
	err := foo()
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT
