package main

import "fmt"

func foo(arg int) (string, error) {

	// ...

	return "something", nil
}

func bar() error {

	val, err := foo(42)
	if err != nil {
		return err
	}
	fmt.Println(val)

	return nil
}
