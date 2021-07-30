package main

import "fmt"

type CustomErr struct {
	Op  string
	Err error
}
func (c *CustomErr) Error() string {
	return fmt.Sprintf("Op: %s, Err: %s", c.Op, c.Err)
}
func foo() error {
	var ce *CustomErr
	return ce
}
func main() {
	err := foo()
	if err != nil {
		fmt.Println("got an error: ", err)
	} else {
		fmt.Println("no error: ", err)
	}
}
