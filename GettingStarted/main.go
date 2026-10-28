package main

import (
	Values "example/hello/01.values"
	Variables "example/hello/02.variables"
	"fmt"
)

func main() {

	fmt.Println("Hello, World!")

	Values.PrintValues()
	Variables.GoVariables()
}
