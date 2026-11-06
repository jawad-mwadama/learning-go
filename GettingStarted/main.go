package main

import (
	values "example/hello/01.values"
	variables "example/hello/02.variables"
	constants "example/hello/03.constants"
	"fmt"
)

func main() {

	fmt.Println("Hello, World!")

	values.PrintValues()
	variables.GoVariables()
	constants.GetConstants()
}
