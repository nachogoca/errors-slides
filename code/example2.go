package main

import (
	"errors"
)

func foo2(arg int) (string, error) {

	return "", errors.New("error")
}
