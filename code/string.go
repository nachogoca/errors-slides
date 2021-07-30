package main

import (
	"errors"
	"fmt"
	"strings"
)

func b() error {
	return errors.New("b not found")
}

func a() {
	err := b()
	// DON'T DO THIS
	if strings.Contains(err.Error(), "not found") {
		fmt.Println("could not find b")
	}
}
